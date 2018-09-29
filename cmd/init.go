package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a go project",
	Long:  `Generates the code for a go project into an empty directory following some template.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flag("name").Value)
		fmt.Println(cmd.Flag("template").Value)
		fmt.Println("init called")

		fmt.Println(isCurrentDirEmpty())
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Defining flags and configuration settings.
	initCmd.Flags().StringP("name", "n", "", "name of the project to be generated (required)")
	initCmd.MarkFlagRequired("name")

	initCmd.Flags().StringP("template", "t", "", "name of template to be used (defaults to hello-world)")
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
