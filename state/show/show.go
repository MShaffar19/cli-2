package show

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bndr/gotabulate"
	"github.com/spf13/cobra"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/constraints"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/internal/updater"
	"github.com/ActiveState/cli/pkg/cmdlets/commands"
	"github.com/ActiveState/cli/pkg/projectfile"
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
	print.Line("%s", project.Name)

	print.BoldInline("%s: ", locale.T("print_state_show_organization"))
	print.Line("%s", project.Owner)

	print.Line("")

	printPlatforms(project)
	printLanguages(project)
	printVariables(project)
	printScripts(project)
	printEvents(project)
}

func printPlatforms(project *projectfile.Project) {
	if len(project.Platforms) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, platform := range project.Platforms {
		constrained := "*"
		if !constraints.PlatformMatches(platform) {
			constrained = ""
		}
		v := fmt.Sprintf("%s%s %s %s (%s)", constrained, platform.Os, platform.Version, platform.Architecture, platform.Name)
		rows = append(rows, []interface{}{v})
	}

	print.BoldInline("%s:", locale.T("print_state_show_platforms"))
	printTable(rows)
}

func printEvents(project *projectfile.Project) {
	if len(project.Events) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, event := range project.Events {
		if !constraints.IsConstrained(event.Constraints) {
			rows = append(rows, []interface{}{event.Name})
		}
	}

	print.BoldInline("%s:", locale.T("print_state_show_events"))
	printTable(rows)
}

func printScripts(project *projectfile.Project) {
	if len(project.Scripts) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, script := range project.Scripts {
		if !constraints.IsConstrained(script.Constraints) {
			rows = append(rows, []interface{}{script.Name, script.Description})
		}
	}

	print.BoldInline("%s:", locale.T("print_state_show_scripts"))
	printTable(rows)
}

func printLanguages(project *projectfile.Project) {
	if len(project.Languages) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, language := range project.Languages {
		if !constraints.IsConstrained(language.Constraints) {
			rows = append(rows, []interface{}{language.Name, language.Version})
		}
	}

	print.BoldInline("%s:", locale.T("print_state_show_languages"))
	printTable(rows)
}

func printVariables(project *projectfile.Project) {
	if len(project.Variables) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, variable := range project.Variables {
		if !constraints.IsConstrained(variable.Constraints) {
			rows = append(rows, []interface{}{variable.Name, variable.Description})
		}
	}

	print.BoldInline("%s:", locale.T("print_state_show_variables"))
	printTable(rows)
}

func printTable(rows [][]interface{}) {
	t := gotabulate.Create(rows)

	// gotabulate tries to make the first row the headers, so use some empty header instead
	// this is also the reason why we're using BoldInLine, since the header line will act as the newline
	t.SetHeaders([]string{""})

	t.SetHideLines([]string{"betweenLine", "top", "aboveTitle", "belowheader", "LineTop", "LineBottom", "bottomLine"}) // Don't print whitespace lines
	t.SetAlign("left")
	print.Line(t.Render("plain"))
}
