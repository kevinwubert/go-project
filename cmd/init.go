package cmd

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/kevinwubert/go-project/pkg/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a go project",
	Long:  `Generates the code for a go project into an empty directory following some template.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := runInit(cmd, args)
		if err != nil {
			log.Errorf("Error initializing go-project: %v", err)
			return
		}
		log.Infof("Successful go-project")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Defining flags and configuration settings.
	initCmd.Flags().StringP("name", "n", "", "name of the project to be generated (required)")
	initCmd.MarkFlagRequired("name")

	initCmd.Flags().StringP("template", "t", "", "name of template to be used (defaults to hello-world)")
}

func runInit(cmd *cobra.Command, args []string) error {
	// isEmpty, err := isCurrentDirEmpty()
	// if err != nil {
	// 	return err
	// }
	// if !isEmpty {
	// 	return errors.New("cannot init go-project in non-empty directory")
	// }

	err := isCurrentDirValid()
	if err != nil {
		return errors.Wrap(err, "cannot init go-project with invalid dir")
	}

	name := cmd.Flag("name").Value.String()
	templateName := cmd.Flag("template").Value.String()

	log.Infof("%v and %v", name, templateName)

	// Ways to get template?
	// Make file to run an initial convertTemplates.go to a static go file
	// so only the binary is needed for the install?
	// Convert templates to a templates.go?
	// I think this way is probably the coolest :)
	// Similar to using packr

	err = templates.Create(templateName, name)
	if err != nil {
		return err
	}

	return nil
}

func isCurrentDirEmpty() (bool, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return false, err
	}
	return isDirEmpty(pwd)
}

// TODO: add viper config to allow ignorable files and
// then generalize this function
func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}

	return false, err
}

func isCurrentDirValid() error {
	pwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "could not get working directory")
	}
	return isDirValid(pwd)
}

func isDirValid(name string) error {
	files, err := ioutil.ReadDir(name)
	if err != nil {
		return errors.Wrap(err, "could not read directory")
	}

	for _, file := range files {
		if !shouldIgnoreFile(file) {
			return errors.New("working directory has unignorable files")
		}
	}

	return nil
}

func shouldIgnoreFile(file os.FileInfo) bool {
	ignoreFiles := []string{".gitignore", ".git", "LICENSE", "README.md"}
	return inStringArr(file.Name(), ignoreFiles)
}

func inStringArr(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
