package activate

import (
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/analytics"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/pkg/cmdlets/commands"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
)

type Activate struct {
	namespaceSelect  namespaceSelectAble
	activateCheckout CheckoutAble
}

type ActivateParams struct {
	Namespace     string
	PreferredPath string
	Output        commands.Output
}

func NewActivate(namespaceSelect namespaceSelectAble, activateCheckout CheckoutAble) *Activate {
	return &Activate{
		namespaceSelect,
		activateCheckout,
	}
}

func (r *Activate) Run(params *ActivateParams) error {
	return r.run(params, activationLoop)
}

func sendProjectIDToAnalytics(namespace string, configFile string) {
	names, fail := project.ParseNamespaceOrConfigfile(namespace, configFile)
	if fail != nil {
		logging.Debug("error resolving namespace: %v", fail.ToError())
		return
	}

	platProject, fail := model.FetchProjectByName(names.Owner, names.Project)
	if fail != nil {
		logging.Debug("error getting platform project: %v", fail.ToError())
		return
	}
	projectID := platProject.ProjectID.String()
	analytics.EventWithLabel(
		analytics.CatBuild, analytics.ActBuildProject, projectID,
	)
}

func (r *Activate) run(params *ActivateParams, activatorLoop activationLoopFunc) error {
	logging.Debug("Activate %v, %v", params.Namespace, params.PreferredPath)

	targetPath, err := r.setupPath(params.Namespace, params.PreferredPath)
	if err != nil {
		return err
	}

	configFile, err := r.setupConfigFile(targetPath)
	if err != nil {
		if params.Namespace == "" {
			logging.Error("Error finding projectfile during activation: %v", err)
			// The default failure returned by the project package is a big too vague,
			// we want to give the user something more actionable for the context they're in
			return failures.FailUserInput.New("err_project_notexist_asyaml")
		}
		err := r.activateCheckout.Run(params.Namespace, targetPath)
		if err != nil {
			return err
		}
	}
	if filepath.Dir(configFile) != targetPath {
		targetPath = filepath.Dir(configFile)
	}

	if params.Output != "" {
		return activateOutput(targetPath, params.Output)
	}

	go sendProjectIDToAnalytics(params.Namespace, configFile)

	return activatorLoop(targetPath, activate)
}

func (r *Activate) setupPath(namespace string, preferredPath string) (string, error) {
	switch {
	// Checkout via namespace (eg. state activate org/project) and set resulting path
	case namespace != "":
		return r.namespaceSelect.Run(namespace, preferredPath)
	// Use the user provided path
	case preferredPath != "":
		return preferredPath, nil
	// Get path from working directory
	default:
		return os.Getwd()
	}
}

func (r *Activate) setupConfigFile(targetPath string) (string, error) {
	proj, err := project.FromPath(targetPath)
	if err != nil {
		return filepath.Join(targetPath, constants.ConfigFileName), err
	}
	return proj.Source().Path(), nil
}
