package virtualenvironment

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	rt "runtime"
	"strings"

	"github.com/google/uuid"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/pkg/platform/runtime"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
)

var persisted *VirtualEnvironment

// FailAlreadyActive is a failure given when a project is already active
var FailAlreadyActive = failures.Type("virtualenvironment.fail.alreadyactive", failures.FailUser)

// OS is used by tests to spoof a different value
var OS = rt.GOOS

// VirtualEnvironment represents our virtual environment, it pulls together and virtualizes the runtime environment
type VirtualEnvironment struct {
	project             *project.Project
	activationID        string
	onDownloadArtifacts func()
	onInstallArtifacts  func()
	onUseCache          func()
	artifactPaths       []string
}

// Get returns a persisted version of VirtualEnvironment{}
func Get() *VirtualEnvironment {
	if persisted == nil {
		persisted = Init()
	}

	return persisted
}

// Init creates an instance of VirtualEnvironment{} with default settings
func Init() *VirtualEnvironment {
	return &VirtualEnvironment{
		project:      project.Get(),
		activationID: uuid.New().String(),
	}
}

func NewWithArtifacts(artifactPaths []string) *VirtualEnvironment {
	return &VirtualEnvironment{
		activationID:  uuid.New().String(),
		artifactPaths: artifactPaths,
	}
}

// Activate the virtual environment
func (v *VirtualEnvironment) Activate() *failures.Failure {
	logging.Debug("Activating Virtual Environment")

	activeProject := os.Getenv(constants.ActivatedStateEnvVarName)
	if activeProject != "" {
		return FailAlreadyActive.New("err_already_active", v.project.Owner()+"/"+v.project.Name())
	}

	if strings.ToLower(os.Getenv(constants.DisableRuntime)) != "true" {
		if failure := v.activateRuntime(); failure != nil {
			return failure
		}
	}

	return nil
}

// OnDownloadArtifacts will call the given function when artifacts are being downloaded
func (v *VirtualEnvironment) OnDownloadArtifacts(f func()) { v.onDownloadArtifacts = f }

// OnInstallArtifacts will call the given function when artifacts are being installed
func (v *VirtualEnvironment) OnInstallArtifacts(f func()) { v.onInstallArtifacts = f }

// OnUseCache will call the given function when the cached runtime is used
func (v *VirtualEnvironment) OnUseCache(f func()) { v.onUseCache = f }

// activateRuntime sets up a runtime environment
func (v *VirtualEnvironment) activateRuntime() *failures.Failure {
	pj := project.Get()
	installer, fail := runtime.NewInstaller(pj.CommitUUID(), pj.Owner(), pj.Name())
	if fail != nil {
		return fail
	}

	installer.OnDownload(v.onDownloadArtifacts)
	installed, fail := installer.Install()
	if fail != nil {
		return fail
	}

	if !installed && v.onUseCache != nil {
		v.onUseCache()
	}

	v.artifactPaths = installer.InstallDirsFromInstallCall()

	return nil
}

// GetEnv returns a map of the cumulative environment variables for all active virtual environments
func (v *VirtualEnvironment) GetEnv(inherit bool, projectDir string) map[string]string {
	env := map[string]string{"PATH": ""}
	if inherit {
		env["PATH"] = os.Getenv("PATH")
	}

	for _, artifactPath := range v.artifactPaths {
		meta, fail := runtime.InitMetaData(artifactPath)
		if fail != nil {
			logging.Warning("Skipping Artifact '%s', could not retrieve metadata: %v", artifactPath, fail)
			continue
		}

		// Unset AffectedEnv
		if meta.AffectedEnv != "" {
			env[meta.AffectedEnv] = ""
		}

		// Set up env according to artifact meta
		templateMeta := struct {
			RelocationDir string
			ProjectDir    string
		}{"", projectDir}
		for k, v := range meta.Env {
			// Dirty workaround until https://www.pivotaltracker.com/story/show/172033094 is implemented
			// This avoids projectDir dependant env vars from being written
			if projectDir == "" && strings.Contains(v, "ProjectDir") {
				continue
			}
			templateMeta.RelocationDir = meta.RelocationDir
			valueTemplate, err := template.New(k).Parse(v)
			if err != nil {
				logging.Error("Skipping artifact with invalid value: %s:%s, error: %v", k, v, err)
				continue
			}
			var realValue bytes.Buffer
			err = valueTemplate.Execute(&realValue, templateMeta)
			if err != nil {
				logging.Error("Skipping artifact whose value could not be parsed: %s:%s, error: %v", k, v, err)
				continue
			}
			env[k] = realValue.String()
		}

		// Set up PATH according to binary locations
		for _, v := range meta.BinaryLocations {
			path := v.Path
			if v.Relative {
				path = filepath.Join(artifactPath, path)
			}
			env["PATH"] = prependPath(env["PATH"], path)
		}

		// Add DLL dir to PATH on Windows
		if meta.RelocationTargetBinaries != "" && rt.GOOS == "windows" {
			env["PATH"] = prependPath(env["PATH"], filepath.Join(meta.Path, meta.RelocationTargetBinaries))
		}
	}

	if projectDir != "" {
		env[constants.ActivatedStateEnvVarName] = projectDir
		env[constants.ActivatedStateIDEnvVarName] = v.activationID
	}

	if inherit {
		return inheritEnv(env)
	}

	return env
}

func prependPath(PATH, prefix string) string {
	var suffix string
	if PATH != "" {
		suffix = string(os.PathListSeparator) + PATH
	}
	return prefix + suffix
}

// GetEnvSlice returns the same results as GetEnv, but formatted in a way that the process package can handle
func (v *VirtualEnvironment) GetEnvSlice(inherit bool) []string {
	envMap := v.GetEnv(inherit, filepath.Dir(projectfile.Get().Path()))
	var env []string
	for k, v := range envMap {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	return env
}

// WorkingDirectory returns the working directory to use for the current environment
func (v *VirtualEnvironment) WorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		// Shouldn't happen unless something is seriously wrong with your system
		panic(locale.T("panic_couldnt_detect_wd", map[string]interface{}{"Error": err.Error()}))
	}

	return wd
}

// ActivationID returns the unique identifier related to the activated instance.
func (v *VirtualEnvironment) ActivationID() string {
	return v.activationID
}
