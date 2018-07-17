package subshell

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) {
	root, err := environment.GetRootPath()
	assert.NoError(t, err, "Should detect root path")
	os.Chdir(filepath.Join(root, "test"))
}

func TestActivate(t *testing.T) {
	setup(t)
	var wg sync.WaitGroup

	os.Setenv("SHELL", "bash")
	os.Setenv("ComSpec", "cmd.exe")
	venv, err := Activate(&wg)

	assert.NoError(t, err, "Should activate")

	assert.NotEqual(t, "", venv.Shell(), "Should detect a shell")
	assert.True(t, venv.IsActive(), "Subshell should be active")

	err = venv.Deactivate()
	assert.NoError(t, err, "Should deactivate")

	assert.False(t, venv.IsActive(), "Subshell should be inactive")
}

func TestActivateFailures(t *testing.T) {
	setup(t)
	var wg sync.WaitGroup

	shell := os.Getenv("SHELL")
<<<<<<< HEAD
	comspec := os.Getenv("ComSpec")

=======
	cmd := os.Getenv("ComSpec")
>>>>>>> master
	os.Setenv("SHELL", "foo")
	os.Setenv("ComSpec", "foo")
	_, err := Activate(&wg)
	os.Setenv("SHELL", shell)
	os.Setenv("ComSpec", cmd)

	os.Setenv("SHELL", shell)
	os.Setenv("ComSpec", comspec)

	assert.Error(t, err, "Should produce an error because of unsupported shell")
}

func TestIsActivated(t *testing.T) {
	assert.False(t, IsActivated(), "Test environment is not in an activated state")
}

func TestRunCommand(t *testing.T) {
	pfile := &projectfile.Project{}
	pfile.Persist()

	os.Setenv("SHELL", "bash")

	subs, err := Get()
	assert.NoError(t, err)

	tmpfile, err := ioutil.TempFile("", "testRunCommand")
	assert.NoError(t, err)
	tmpfile.Close()
	os.Remove(tmpfile.Name())

	if runtime.GOOS != "windows" {
		subs.Run(fmt.Sprintf(`echo "Hello"
touch %s`, tmpfile.Name()))
	} else {
		subs.Run(fmt.Sprintf(`echo "Hello"
copy NUL %s`, tmpfile.Name()))
	}

	assert.FileExists(t, tmpfile.Name())
}
