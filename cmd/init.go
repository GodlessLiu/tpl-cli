package cmd

import "github.com/spf13/cobra"

var InitTemplateCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project with the template",
	Long:  "Initialize the project with the template",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
