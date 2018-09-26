package cmd

import (
	"fmt"

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
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Defining flags and configuration settings.
	initCmd.Flags().StringP("name", "n", "", "name of the project to be generated (required)")
	initCmd.MarkFlagRequired("name")

	initCmd.Flags().StringP("template", "t", "", "name of template to be used (defaults to hello-world)")
}
