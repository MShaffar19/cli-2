package show

import (
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/constraints"
	"github.com/ActiveState/cli/internal/expander"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/internal/scm"
	"github.com/ActiveState/cli/internal/updater"
	"github.com/ActiveState/cli/pkg/cmdlets/commands"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/spf13/cobra"
)

// Command is the show command's definition.
var Command = &commands.Command{
	Name:        "show",
	Description: "show_project",
	Run:         Execute,

	Arguments: []*commands.Argument{
		&commands.Argument{
			Name:        "remote",
			Description: "arg_state_show_remote_description",
			Variable:    &Args.Remote,
		},
	},
}

// Args holds the arg values passed through the command line.
var Args struct {
	Remote string
}

// Execute the show command.
func Execute(cmd *cobra.Command, args []string) {
	logging.Debug("Execute")

	updater.PrintUpdateMessage()

	var project *projectfile.Project
	if Args.Remote == "" {
		project = projectfile.Get()
	} else if scm := scm.FromRemote(Args.Remote); scm != nil {
		// TODO: remote fetching of activestate.yaml and parsing
	} else {
		path := Args.Remote
		projectFile := filepath.Join(Args.Remote, constants.ConfigFileName)
		if _, err := os.Stat(path); err != nil {
			print.Error(locale.T("err_state_show_path_does_not_exist"))
			return
		} else if _, err := os.Stat(projectFile); err != nil {
			print.Error(locale.T("err_state_show_no_config"))
			return
		}
		var err error
		project, err = projectfile.Parse(projectFile)
		if err != nil {
			logging.Errorf("Unable to parse activestate.yaml: %s", err)
			print.Error(locale.T("err_state_show_project_parse"))
			return
		}
	}

	print.BoldInline("%s: ", locale.T("print_state_show_name"))
	print.Formatted("%s\n", project.Name)

	print.BoldInline("%s: ", locale.T("print_state_show_organization"))
	print.Formatted("%s\n", project.Owner)

	//print.Bold("%s: \n", locale.T("print_state_show_url"))
	//print.Formatted("%s\n", "")

	if len(project.Platforms) > 0 {
		print.Bold("%s:", locale.T("print_state_show_platforms"))
		for _, platform := range project.Platforms {
			constrained := "*"
			if !constraints.PlatformMatches(platform) {
				constrained = " "
			}
			print.Formatted(" %s%s %s %s (%s)\n", constrained, platform.Os, platform.Version, platform.Architecture, platform.Name)
		}
	}

	if len(project.Events) > 0 {
		print.Bold("%s:", locale.T("print_state_show_events"))
		for _, event := range project.Events {
			if !constraints.IsConstrained(event.Constraints) {
				value := expander.ExpandFromProject(event.Value, project)
				print.Formatted("  %s: %s\n", event.Name, value)
			}
		}
	}

	if len(project.Scripts) > 0 {
		print.Bold("%s:", locale.T("print_state_show_scripts"))
		for _, script := range project.Scripts {
			if !constraints.IsConstrained(script.Constraints) {
				value := expander.ExpandFromProject(script.Value, project)
				print.Formatted("  %s: %s\n", script.Name, value)
			}
		}
	}

	if len(project.Languages) > 0 {
		print.Bold("%s:", locale.T("print_state_show_languages"))
		for _, language := range project.Languages {
			if !constraints.IsConstrained(language.Constraints) {
				print.Formatted("  %s %s (%d %s)\n", language.Name, language.Version, len(language.Packages), locale.T("print_state_show_packages"))
			}
		}
	}

	if len(project.Variables) > 0 {
		print.Bold("%s:", locale.T("print_state_show_env_vars"))
		for _, variable := range project.Variables {
			if !constraints.IsConstrained(variable.Constraints) {
				print.Formatted(" - %s\n", variable.Name)
			}
		}
	}
}
