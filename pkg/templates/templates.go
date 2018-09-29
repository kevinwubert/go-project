package templates

// DefaultTemplatesDir is the default name for the templates dir to be stored
const DefaultTemplatesDir = "/templates"

// Create either creates the template based off the template name
// and templated by name or returns an error if that template does not exist
func Create(templateName string, name string) error {
	return nil
}

// Recursive struct
type directory struct {
	name  string
	dirs  []*directory
	files []file
}

type file struct {
	name     string
	contents []byte
}

type template struct {
	name    string
	rootDir directory
}

// ProcessTemplatesDir takes in templates
func ProcessTemplatesDir(dir string) {

}
