package templates

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

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
	if ts != nil {
		for _, t := range ts {
			s += t.CodeString()
		}
	}

	s += `}
`

	return s
}

func (t template) CodeString() string {
	s := `"` + t.name + `": template{
name: "` + t.name + `",
rootDir:
`

	s += t.rootDir.CodeString()
	s += `},
`
	return s
}

func (d directory) CodeString() string {
	s := `&directory{
name:  "` + d.name + `",
files: []file{
`
	for _, f := range d.files {
		s += f.CodeString()
	}

	s += `
},
dirs:  []*directory{
`
	for _, internalDir := range d.dirs {
		s += internalDir.CodeString()
	}

	s += `
},
},
`
	return s
}

func (f file) CodeString() string {
	bytesStr := fmt.Sprintf("%v", f.data)
	bytesStr = strings.Replace(bytesStr, " ", ", ", -1)

	s := `file{
name: "` + f.name + `",
data: []byte{` + bytesStr[1:len(bytesStr)-1] + `},
},`
	return s
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

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			t, err := buildTemplateFromDir(dir+"/"+f.Name(), f.Name())
			if err != nil {
				return err
			}
			ts[f.Name()] = t
		}
	}

	err = util.CreateFile(filename, []byte(ts.CodeString()))
	return err
}

func buildTemplateFromDir(path string, name string) (template, error) {
	rootDir, err := buildDirFromDir(path, "")
	if err != nil {
		return template{}, nil
	}

	t := template{
		name:    name,
		rootDir: rootDir,
	}

	return t, nil
}

func buildDirFromDir(dirpath string, name string) (*directory, error) {
	ds := []*directory{}
	fs := []file{}

	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			d, err := buildDirFromDir(dirpath+"/"+file.Name(), file.Name())
			if err != nil {
				return nil, err
			}
			ds = append(ds, d)
		} else {
			f, err := buildFileFromFile(dirpath+"/"+file.Name(), file.Name())
			if err != nil {
				return nil, err
			}
			fs = append(fs, f)
		}
	}

	dir := &directory{
		name:  name,
		files: fs,
		dirs:  ds,
	}

	return dir, nil
}

func buildFileFromFile(path string, name string) (file, error) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return file{}, err
	}
	f := file{
		name: name,
		data: input,
	}
	return f, nil
}
