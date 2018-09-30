package templates

import (
	"errors"

	"github.com/kevinwubert/go-project/pkg/util"
)

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
	rootDir *directory
}

type templates map[string]template

func (t template) Generate() error {
	err := t.rootDir.Create(".")
	return err
}

func (d directory) Create(path string) error {
	currPath := path + "/" + d.name

	if currPath != "./" {
		err := util.CreateDir(currPath)
		if err != nil {
			return err
		}
	}

	for _, f := range d.files {
		err := f.Create(currPath)
		if err != nil {
			return err
		}
	}
	for _, internalDir := range d.dirs {
		err := internalDir.Create(currPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f file) Create(path string) error {
	filepath := path + "/" + f.name
	err := util.CreateFile(filepath, f.data)
	if err != nil {
		return err
	}
	return nil
}

func (ts templates) CodeString() string {
	s := `package templates

var staticTemplates = templates{
`

	for _, t := range ts {
		s += t.CodeString()
	}

	s += `}`

	return s
}

func (t template) CodeString() string {
	s := `"hello-world": template{
name: "hello-world",
rootDir:`

	s += t.rootDir.CodeString()
	s += `
	},`
	return s
}

func (d directory) CodeString() string {
	return "blah"
}

func (f file) CodeString() string {
	return "blah"
}

// Create either creates the template based off the template name
// and templated by name or returns an error if that template does not exist
func Create(templateName string, name string) error {
	if _, ok := staticTemplates[templateName]; !ok {
		return errors.New("static template " + templateName + " not found")
	}
	return staticTemplates[templateName].Generate()
}

// DefaultTemplatesDir is the default name for the templates dir to be stored
const DefaultTemplatesDir = "./templates"

// ProcessTemplatesDir creates a statictemplates.go from the directory dir
// statictemplates contains a slice of template
func ProcessTemplatesDir(dir string) error {
	filename := "./pkg/templates/static_templates.go"
	ts := make(templates)

	// t := template{
	// 	name: "stuff",
	// 	rootDir: directory{
	// 		name:  "thing",
	// 		files: []file{},
	// 		dirs:  []*directory{},
	// 	},
	// }
	t, err := buildTemplateFromDir(dir)
	if err != nil {
		return err
	}
	ts[t.name] = t

	err = util.CreateFile(filename, []byte(ts.CodeString()))

	return err
}

func buildTemplateFromDir(dir string) (template, error) {
	return template{}, nil
}
