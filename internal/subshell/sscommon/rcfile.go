package sscommon

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/mash/go-tempfile-suffix"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/internal/virtualenvironment"
	"github.com/ActiveState/cli/pkg/project"
)

func WriteRcFile(rcTemplateName string, path string, env map[string]string) *failures.Failure {
	if fail := fileutils.Touch(path); fail != nil {
		return fail
	}

	if fail := cleanRcFile(path); fail != nil {
		return fail
	}

	box := packr.NewBox("../../../assets/shells")
	tpl := box.String(rcTemplateName)

	rcData := map[string]interface{}{
		"Start": constants.RCAppendStartLine,
		"Stop":  constants.RCAppendStopLine,
		"Env":   env,
	}
	t, err := template.New("rcfile_append").Parse(tpl)
	if err != nil {
		return failures.FailTemplating.Wrap(err)
	}

	var out bytes.Buffer
	err = t.Execute(&out, rcData)
	if err != nil {
		return failures.FailTemplating.Wrap(err)
	}

	logging.Debug("Writing to %s:\n%s", path, out.String())

	return fileutils.AppendToFile(path, []byte(fileutils.LineEnd+out.String()))
}

func cleanRcFile(path string) *failures.Failure {
	readFile, err := os.Open(path)

	if err != nil {
		return failures.FailIO.Wrap(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	var strip bool
	var fileContents []string
	for scanner.Scan() {
		text := scanner.Text()

		// Detect start line
		if strings.Contains(text, constants.RCAppendStartLine) {
			logging.Debug("Cleaning previous RC lines from %s", path)
			strip = true
		}

		// Strip line
		if strip {
			continue
		}

		// Rebuild file contents
		fileContents = append(fileContents, scanner.Text())

		// Detect stop line
		if strings.Contains(text, constants.RCAppendStopLine) {
			strip = false
		}
	}
	readFile.Close()

	return fileutils.WriteFile(path, []byte(strings.Join(fileContents, fileutils.LineEnd)))
}

// SetupShellRcFile create a rc file to activate a runtime (without a project being present)
func SetupShellRcFile(rcFileName, templateName string, env map[string]string, namespace project.Namespaced) error {
	box := packr.NewBox("../../../assets/shells")
	tpl := box.String(templateName)

	rcData := map[string]interface{}{
		"Env":     env,
		"Project": namespace.String(),
	}
	t, err := template.New("rcfile").Parse(tpl)
	if err != nil {
		return errs.Wrap(err, "Failed to parse template file.")
	}

	var out bytes.Buffer
	err = t.Execute(&out, rcData)
	if err != nil {
		return errs.Wrap(err, "failed to execute template.")
	}

	f, err := os.Create(rcFileName)
	if err != nil {
		return locale.WrapError(err, "sscommon_rc_file_creation_err", "Failed to create file {{.V0}}", rcFileName)
	}
	defer f.Close()

	f.WriteString(out.String())

	err = os.Chmod(rcFileName, 0755)
	if err != nil {
		return errs.Wrap(err, "Failed to set executable flag.")
	}
	return nil
}

// SetupProjectRcFile creates a temporary RC file that our shell is initiated from, this allows us to template the logic
// used for initialising the subshell
func SetupProjectRcFile(templateName, ext string, env map[string]string) (*os.File, *failures.Failure) {
	box := packr.NewBox("../../../assets/shells")
	tpl := box.String(templateName)
	prj := project.Get()

	userScripts := ""
	for _, event := range prj.Events() {
		if event.Name() == "ACTIVATE" {
			userScripts = userScripts + "\n" + event.Value()
		}
	}

	inuse := []string{}
	scripts := map[string]string{}
	var explicitName string

	// Prepare script map to be parsed by template
	for _, cmd := range prj.Scripts() {
		explicitName = fmt.Sprintf("%s_%s", prj.NormalizedName(), cmd.Name())

		_, err := exec.LookPath(cmd.Name())
		if err == nil {
			// Do not overwrite commands that are already in use and
			// keep track of those commands to warn to the user
			inuse = append(inuse, cmd.Name())
			continue
		}

		scripts[cmd.Name()] = cmd.Name()
		scripts[explicitName] = cmd.Name()
	}

	if len(inuse) > 0 {
		print.Warning(locale.Tr("warn_script_name_in_use", strings.Join(inuse, "\n  - "), inuse[0], prj.NormalizedName(), explicitName))
	}

	rcData := map[string]interface{}{
		"Owner":       prj.Owner(),
		"Name":        prj.Name(),
		"Env":         env,
		"WD":          virtualenvironment.Get().WorkingDirectory(),
		"UserScripts": userScripts,
		"Scripts":     scripts,
	}
	t, err := template.New("rcfile").Parse(tpl)
	if err != nil {
		return nil, failures.FailTemplating.Wrap(err)
	}

	var out bytes.Buffer
	err = t.Execute(&out, rcData)
	if err != nil {
		return nil, failures.FailTemplating.Wrap(err)
	}

	tmpFile, err := tempfile.TempFileWithSuffix(os.TempDir(), "state-subshell-rc", ext)
	if err != nil {
		return nil, failures.FailOS.Wrap(err)
	}
	defer tmpFile.Close()

	tmpFile.WriteString(out.String())

	logging.Debug("Using project RC: %s", out.String())

	return tmpFile, nil
}
