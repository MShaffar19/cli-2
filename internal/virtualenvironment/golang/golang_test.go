package golang

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"

	"github.com/ActiveState/ActiveState-CLI/internal/artifact"
	"github.com/ActiveState/ActiveState-CLI/internal/config"
	"github.com/ActiveState/ActiveState-CLI/internal/constants"
	"github.com/ActiveState/ActiveState-CLI/internal/distribution"
	"github.com/ActiveState/ActiveState-CLI/internal/environment"
	"github.com/ActiveState/ActiveState-CLI/pkg/projectfile"
)

func setup(t *testing.T) {
	root, _ := environment.GetRootPath()
	os.Chdir(filepath.Join(root, "test"))
}

func TestLanguage(t *testing.T) {
	venv := &VirtualEnvironment{}
	assert.Equal(t, "go", venv.Language(), "Should return go")
}

func TestDataDir(t *testing.T) {
	venv := &VirtualEnvironment{}
	assert.Empty(t, venv.DataDir())

	venv.SetDataDir("/foo")
	assert.NotEmpty(t, venv.DataDir(), "Should set the datadir")
}

func TestLanguageMeta(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}
	assert.Nil(t, venv.Artifact(), "Should not have artifact info")

	venv.SetArtifact(&artifact.Artifact{
		Meta: &artifact.Meta{
			Name: "test",
		},
		Path: "test",
	})
	assert.NotNil(t, venv.Artifact(), "Should have artifact info")
}

func TestLoadPackageFromPath(t *testing.T) {
	venv := &VirtualEnvironment{}

	datadir := filepath.Join(os.TempDir(), "as-state-test")
	os.RemoveAll(datadir)
	os.Mkdir(datadir, os.ModePerm)
	venv.SetDataDir(datadir)

	dist, fail := distribution.Obtain()
	assert.NoError(t, fail.ToError())

	var language *artifact.Artifact
	for _, lang := range dist.Languages {
		if strings.ToLower(lang.Meta.Name) == venv.Language() {
			language = lang
			break
		}
	}

	assert.NotNil(t, language, "Should retrieve language from dist")

	artf := dist.Artifacts[dist.Languages[0].Hash][0]
	fail = venv.LoadArtifact(artf)
	assert.NoError(t, fail.ToError(), "Loads artifact without errors")

	// Todo: Test with datadir as source, not the archived version
	assert.FileExists(t, filepath.Join(datadir, "src", artf.Meta.Name, "artifact.json"), "Should create a package symlink")
}

func TestActivate(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}

	venv.SetArtifact(&artifact.Artifact{
		Meta: &artifact.Meta{
			Name:    "go",
			Version: "1.10",
		},
		Path: "test",
	})

	datadir := config.GetDataDir()
	datadir = filepath.Join(datadir, "test")
	venv.SetDataDir(datadir)

	venv.Activate()

	assert.DirExists(t, filepath.Join(venv.DataDir(), "bin"))
}

func TestNamespace(t *testing.T) {
	setup(t)

	venv := &VirtualEnvironment{}
	root, err := environment.GetRootPath()
	assert.NoError(t, err, "Should get root path")

	project := projectfile.Project{}
	dat := strings.TrimSpace(`
name: Bar
owner: Foo`)

	err = yaml.Unmarshal([]byte(dat), &project)
	assert.NoError(t, err, "Should create project struct")
	project.SetPath(filepath.Join(root, "foo"))
	project.Persist()

	ns := venv.namespace()
	assert.Equal(t, constants.LibraryNamespace+constants.LibraryName, ns, "Should detect namespace from remote")

	project = projectfile.Project{}
	dat = strings.TrimSpace(`
name: Bar
owner: Foo
namespace: foo.bar/foo/bar`)

	err = yaml.Unmarshal([]byte(dat), &project)
	assert.NoError(t, err, "Should create project struct")
	project.Persist()

	ns = venv.namespace()
	assert.Equal(t, "foo.bar/foo/bar", ns, "Should detect namespace from namespace property")

	project = projectfile.Project{}
	dat = strings.TrimSpace(`
name: Bar
owner: Foo`)

	err = yaml.Unmarshal([]byte(dat), &project)
	assert.NoError(t, err, "Should create project struct")
	project.Persist()

	ns = venv.namespace()
	assert.Equal(t, path.Join(constants.DefaultNamespaceDomain, project.Owner, project.Name), ns, "Should create namespace based on owner/name")
}
