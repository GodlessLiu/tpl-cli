package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(VersionCmd)
	rootCmd.AddCommand(InitTemplateCmd)
}

var rootCmd = &cobra.Command{
	Use:   "tpl",
	Short: "tpl is a tool for generating code",
	Long:  "tpl is a tool generate code based on templates, powered by go",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tpl is a tool for generating code")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
