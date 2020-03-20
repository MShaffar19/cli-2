package runtime

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/progress"
	"github.com/ActiveState/cli/internal/unarchiver"
	"github.com/ActiveState/cli/pkg/platform/runtime/envdef"
	"github.com/go-openapi/strfmt"
)

var _ Assembler = &AlternativeRuntime{}

// AlternativeRuntime holds all the meta-data necessary to activate a runtime
// environment for an Alternative build
type AlternativeRuntime struct {
	cacheDir       string
	recipeID       strfmt.UUID
	artifactMap    map[string]*HeadChefArtifact
	artifactOrder  map[string]int
	tempInstallDir string
}

// NewAlternativeRuntime returns a new alternative runtime assembler
// It filters the provided artifact list for useable artifacts
// The recipeID is needed to define the installation directory
func NewAlternativeRuntime(artifacts []*HeadChefArtifact, cacheDir string, recipeID strfmt.UUID) (*AlternativeRuntime, *failures.Failure) {

	artifactMap := map[string]*HeadChefArtifact{}
	artifactOrder := map[string]int{}

	ar := &AlternativeRuntime{
		cacheDir: cacheDir,
		recipeID: recipeID,
	}
	for i, artf := range artifacts {

		if artf.URI == "" {
			continue
		}
		filename := filepath.Base(artf.URI.String())
		if !strings.HasSuffix(filename, ar.InstallerExtension()) {
			continue
		}

		// XXX: For now we are excluding terminal artifacts ie., the artifacts that a packaging step would produce.
		// Right now, these artifacts are empty anyways...
		if artf.IngredientVersionID == "" {
			continue
		}
		downloadDir := ar.downloadDirectory(artf)

		artifactMap[downloadDir] = artf
		artifactOrder[artf.ArtifactID.String()] = i
	}

	if len(artifactMap) == 0 {
		return ar, FailNoValidArtifact.New(locale.T("err_no_valid_artifact"))
	}

	ar.artifactMap = artifactMap
	ar.artifactOrder = artifactOrder
	return ar, nil
}

// InstallerExtension is always .tar.gz
func (ar *AlternativeRuntime) InstallerExtension() string {
	return ".tar.gz"
}

// Unarchiver always returns an unarchiver for gzipped tarballs
func (ar *AlternativeRuntime) Unarchiver() unarchiver.Unarchiver {
	return unarchiver.NewTarGz()
}

// BuildEngine always returns Alternative
func (ar *AlternativeRuntime) BuildEngine() BuildEngine {
	return Alternative
}

func (ar *AlternativeRuntime) cachedArtifact(downloadDir string) *string {
	candidate := filepath.Join(downloadDir, constants.ArtifactArchiveName)
	if fileutils.FileExists(candidate) {
		return &candidate
	}

	return nil
}

// ArtifactsToDownloadAndUnpack returns the artifacts that we need to download
// for this project.
// It returns nothing if the final installation directory is non-empty.
// It filters out artifacts that have been downloaded before, and adds them to
// the list of artifacts that need to be unpacked only.
func (ar *AlternativeRuntime) ArtifactsToDownloadAndUnpack() ([]*HeadChefArtifact, map[string]*HeadChefArtifact) {
	downloadArtfs := []*HeadChefArtifact{}
	archives := map[string]*HeadChefArtifact{}

	// if final installation directory exists -> no need to download or unpack anything
	if fileutils.DirExists(ar.finalInstallationDirectory()) {
		return downloadArtfs, archives
	}

	for downloadDir, artf := range ar.artifactMap {
		cached := ar.cachedArtifact(downloadDir)
		if cached == nil {
			downloadArtfs = append(downloadArtfs, artf)
		} else {
			archives[*cached] = artf
		}
	}
	return downloadArtfs, archives
}

func (ar *AlternativeRuntime) downloadDirectory(artf *HeadChefArtifact) string {
	return filepath.Join(ar.cacheDir, "artifacts", shortHash(artf.ArtifactID.String()))
}

// DownloadDirectory returns the local directory where the artifact files should
// be downloaded to
func (ar *AlternativeRuntime) DownloadDirectory(artf *HeadChefArtifact) (string, *failures.Failure) {
	p := ar.downloadDirectory(artf)
	fail := fileutils.MkdirUnlessExists(p)
	return p, fail
}

func (ar *AlternativeRuntime) finalInstallationDirectory() string {
	finstDir := filepath.Join(ar.cacheDir, shortHash(ar.recipeID.String()))
	return finstDir
}

// InstallationDirectory returns the local directory where the artifact files
// should be unpacked to.
// For alternative build artifacts, all artifacts are unpacked into the same
// directory eventually.
func (ar *AlternativeRuntime) InstallationDirectory(artf *HeadChefArtifact) string {
	return ar.finalInstallationDirectory()
}

// PreInstall ensures that the final installation directory exists, and is useable
func (ar *AlternativeRuntime) PreInstall() *failures.Failure {
	installDir := ar.finalInstallationDirectory()

	if fileutils.FileExists(installDir) {
		// install-dir exists, but is a regular file
		return FailInstallDirInvalid.New("installer_err_installdir_isfile", installDir)
	}

	if fail := fileutils.MkdirUnlessExists(installDir); fail != nil {
		return fail
	}

	if isEmpty, fail := fileutils.IsEmptyDir(installDir); fail != nil || !isEmpty {
		if fail != nil {
			return fail
		}
		return FailInstallDirInvalid.New("installer_err_installdir_notempty", installDir)
	}

	return nil
}

// PreUnpackArtifact does nothing
func (ar *AlternativeRuntime) PreUnpackArtifact(artf *HeadChefArtifact) *failures.Failure {
	return nil
}

// PostUnpackArtifact is called after the artifacts are unpacked
// In this steps, the artifact contents are moved to its final destination.
// This step also sets up the runtime environment variables.
func (ar *AlternativeRuntime) PostUnpackArtifact(artf *HeadChefArtifact, tmpRuntimeDir string, archivePath string, counter progress.Incrementer) *failures.Failure {

	// final installation target
	ft := ar.InstallationDirectory(artf)

	logging.Debug("ft=%s trd=%s\n", ft, tmpRuntimeDir)

	rt, err := envdef.NewEnvironmentDefinition(filepath.Join(tmpRuntimeDir, constants.RuntimeDefinitionFilename))
	if err != nil {
		return failures.FailRuntime.Wrap(err)
	}
	rt = rt.ReplaceInstallDir(ft)

	// move files to the final installation directory
	fail := fileutils.MoveAllFilesRecursively(
		filepath.Join(tmpRuntimeDir, rt.InstallDir),
		ft, func() { counter.Increment() })
	if fail != nil {
		return fail
	}

	// move the runtime.json to the runtime environment directory
	artifactIndex, ok := ar.artifactOrder[artf.ArtifactID.String()]
	if !ok {
		return failures.FailRuntime.New(fmt.Sprintf("Could not write runtime.json: artifact order for %s unknown", artf.ArtifactID.String()))
	}

	fail = fileutils.MkdirUnlessExists(ar.runtimeEnvBaseDir())
	if fail != nil {
		return fail
	}

	err = rt.WriteFile(filepath.Join(ar.runtimeEnvBaseDir(), fmt.Sprintf("%06d.json", artifactIndex)))
	if err != nil {
		return failures.FailRuntime.Wrap(err, "Failed to write runtime.json to final installation directory %s", ar.runtimeEnvBaseDir())
	}

	if err := os.RemoveAll(tmpRuntimeDir); err != nil {
		logging.Error("removing %s after unpacking runtime: %v", tmpRuntimeDir, err)
		return FailRuntimeInstallation.New("installer_err_runtime_rm_installdir", tmpRuntimeDir)
	}
	return nil
}

func (ar *AlternativeRuntime) runtimeEnvBaseDir() string {
	return filepath.Join(ar.finalInstallationDirectory(), constants.LocalRuntimeEnvironmentDirectory)
}

func (ar *AlternativeRuntime) getRuntimeDefinition() (*envdef.EnvironmentDefinition, error) {

	mergedRuntimeDefinitionFile := filepath.Join(ar.runtimeEnvBaseDir(), constants.RuntimeDefinitionFilename)
	if fileutils.FileExists(mergedRuntimeDefinitionFile) {
		rt, err := envdef.NewEnvironmentDefinition(mergedRuntimeDefinitionFile)
		if err == nil {
			return rt, nil
		}
		logging.Warning("Failed to unmarshal the merged runtime definition file at %s", mergedRuntimeDefinitionFile)
	}

	files, err := ioutil.ReadDir(ar.runtimeEnvBaseDir())
	if err != nil {
		logging.Warning("no environment definition files found")
		return nil, err
	}

	filenames := make([]string, 0, len(files))
	for _, f := range files {
		if ok, _ := regexp.MatchString("[0-9]{6}.json", f.Name()); ok {
			filenames = append(filenames, f.Name())
		}
	}
	sort.Strings(filenames)

	var rtGlobal *envdef.EnvironmentDefinition

	for _, fn := range filenames {
		rtPath := filepath.Join(ar.runtimeEnvBaseDir(), fn)
		rt, err := envdef.NewEnvironmentDefinition(rtPath)
		if err != nil {
			logging.Warning("Failed to read environment definition file %s", rtPath)
			continue
		}
		if rtGlobal == nil {
			rtGlobal = rt
			continue
		}
		rtGlobal, err = rtGlobal.Merge(rt)
		if err != nil {
			logging.Warning("Failed to merge environment definition file %s: %v", rtPath, err)
			continue
		}
	}

	if rtGlobal == nil {
		return nil, errors.New("did not find any runtime definition files")
	}

	err = rtGlobal.WriteFile(mergedRuntimeDefinitionFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to write merged runtime definition file at %s", mergedRuntimeDefinitionFile)
	}

	return rtGlobal, nil
}

// GetEnv returns the environment variable configuration for this build
func (ar *AlternativeRuntime) GetEnv() map[string]string {

	rt, err := ar.getRuntimeDefinition()
	if err != nil {
		logging.Warning("No runtime definition found")
		return map[string]string{}
	}
	return rt.GetEnv()
}
