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
	Create(path string) error
}

func (t template) Generate() error {
	err := t.rootDir.Create("")
	return err
}

func (d directory) Create(path string) error {
	currPath := path + "/" + d.name

	if currPath != "/" {
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

// Create either creates the template based off the template name
// and templated by name or returns an error if that template does not exist
func Create(templateName string, name string) error {
	dir3 := directory{
		name: "zxcvz",
		files: []file{
			file{
				name: "zxcv.txt",
				data: []byte("zxcvzxcv\n"),
			},
			file{
				name: "zxcvzxcv.txt",
				data: []byte("zxcvzxcvzxcvzxcvzxcv\n"),
			},
		},
		dirs: []*directory{},
	}
	dir2 := directory{
		name: "asdfasdf",
		files: []file{
			file{
				name: "asdfffff.txt",
				data: []byte("asdfasdfasdfasdf\n"),
			},
			file{
				name: "asdffssssfff.txt",
				data: []byte("asdfasdfasdfasdf\n"),
			},
		},
		dirs: []*directory{},
	}
	dir := directory{
		name: "test",
		files: []file{
			file{
				name: "test.txt",
				data: []byte("testing123\n"),
			},
		},
		dirs: []*directory{&dir2},
	}
	rdir := directory{
		name: "",
		files: []file{
			file{
				name: "root.txt",
				data: []byte("config\n"),
			},
			file{
				name: "empty.txt",
				data: []byte("config\n"),
			},
		},
		dirs: []*directory{&dir, &dir3},
	}

	return rdir.Create(".")
}

// DefaultTemplatesDir is the default name for the templates dir to be stored
const DefaultTemplatesDir = "/templates"

// ProcessTemplatesDir takes in templates
func ProcessTemplatesDir(dir string) {

}
