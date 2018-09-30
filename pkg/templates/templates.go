package templates

import "github.com/kevinwubert/go-project/pkg/util"

// Recursive struct
type directory struct {
	name  string
	dirs  []*directory
	files []file
}

type file struct {
	name string
	data []byte
}

type template struct {
	name    string
	rootDir directory
}

// Creator has a create function attached to it
type Creator interface {
	Create() error
}

func (t template) Create() error {
	err := t.rootDir.Create()
	return err
}

func (d directory) Create() error {
	for _, f := range d.files {
		err := f.Create()
		if err != nil {
			return err
		}
	}
	for _, internalDir := range d.dirs {
		err := internalDir.Create()
		if err != nil {
			return err
		}
	}
	return nil
}

func (f file) Create() error {
	err := util.CreateFile(f.name, f.data)
	if err != nil {
		return err
	}
	return nil
}

// Create either creates the template based off the template name
// and templated by name or returns an error if that template does not exist
func Create(templateName string, name string) error {
	f := file{
		name: "test.txt",
		data: []byte("testing123\n"),
	}

	return f.Create()
}

// DefaultTemplatesDir is the default name for the templates dir to be stored
const DefaultTemplatesDir = "/templates"

// ProcessTemplatesDir takes in templates
func ProcessTemplatesDir(dir string) {

}
