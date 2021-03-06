package virtualenvironment

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/locale"
	rtmock "github.com/ActiveState/cli/pkg/platform/runtime/mock"
	"github.com/ActiveState/cli/pkg/projectfile"
)

var rtMock *rtmock.Mock

func setup(t *testing.T) {
	root, err := environment.GetRootPath()
	assert.NoError(t, err, "Should detect root path")

	err = os.Chdir(filepath.Join(root, "internal", "virtualenvironment", "testdata"))
	assert.NoError(t, err, "unable to chdir to testdata dir")

	rtMock = rtmock.Init()
	rtMock.MockFullRuntime()

	os.Unsetenv(constants.ActivatedStateEnvVarName)
	os.Unsetenv(constants.ActivatedStateIDEnvVarName)
}

func teardown() {
	projectfile.Reset()
	rtMock.Close()
}

func TestPersist(t *testing.T) {
	setup(t)
	defer teardown()

	v1 := Get()
	v2 := Get()
	assert.True(t, v1 == v2, "Should return same pointer")
}

func TestEvents(t *testing.T) {
	venv := Init()
	onDownloadCalled := false
	onInstallCalled := false

	venv.OnDownloadArtifacts(func() { onDownloadCalled = true })
	venv.OnInstallArtifacts(func() { onInstallCalled = true })

	venv.onDownloadArtifacts()
	venv.onInstallArtifacts()

	assert.True(t, onDownloadCalled, "OnDownloadArtifacts is triggered")
	assert.True(t, onInstallCalled, "OnInstallArtifacts is triggered")
}

func TestActivateFailureUnknownLanguage(t *testing.T) {
	setup(t)
	defer teardown()

	os.Setenv(constants.DisableRuntime, "false")
	defer os.Unsetenv(constants.DisableRuntime)

	project := projectfile.Get()
	project.Languages = append(project.Languages, projectfile.Language{
		Name: "foo",
	})
	project.Persist()

	venv := Init()
	err := venv.Activate()
	assert.Error(t, err, "Should not activate due to unknown language")
}

func TestActivateFailureAlreadyActive(t *testing.T) {
	setup(t)
	defer teardown()

	os.Setenv(constants.ActivatedStateEnvVarName, "test")

	venv := Init()
	failure := venv.Activate()
	namespace := venv.project.Owner() + "/" + venv.project.Name()
	require.NotNil(t, failure, "expected a failure")
	assert.Equal(t, FailAlreadyActive, failure.Type)
	assert.Equal(t, locale.Tr("err_already_active", namespace), failure.Error())
}

func TestEnv(t *testing.T) {
	setup(t)
	defer teardown()

	os.Setenv(constants.DisableRuntime, "true")
	defer os.Unsetenv(constants.DisableRuntime)

	os.Setenv(constants.ProjectEnvVarName, projectfile.Get().Path())
	defer os.Unsetenv(constants.ProjectEnvVarName)

	venv := Init()
	env, err := venv.GetEnv(false, filepath.Dir(projectfile.Get().Path()))
	require.NoError(t, err)

	assert.NotContains(t, env, constants.ProjectEnvVarName)
	assert.NotEmpty(t, env[constants.ActivatedStateIDEnvVarName])
	assert.NotEmpty(t, venv.ActivationID())
}

func TestInheritEnv_MultipleEquals(t *testing.T) {
	key := "MULTIPLEEQUALS"
	value := "one=two two=three three=four"

	os.Setenv(key, value)
	defer os.Unsetenv(key)

	env := map[string]string{}
	updated := inheritEnv(env)

	assert.Equal(t, value, updated[key])
}

func TestSkipActivateRuntimeEnvironment(t *testing.T) {
	setup(t)
	defer teardown()

	os.Setenv(constants.DisableRuntime, "true")
	defer os.Unsetenv(constants.DisableRuntime)

	project := projectfile.Project{}
	dat := strings.TrimSpace(`
project: "https://platform.activestate.com/string/string?commitID=00010001-0001-0001-0001-000100010001"
languages:
    - name: Python3`)
	yaml.Unmarshal([]byte(dat), &project)
	project.Persist()

	venv := Init()
	fail := venv.Activate()
	require.NoError(t, fail.ToError(), "Should activate")
}
