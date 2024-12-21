package cmd

import (
	"fmt"

	"github.com/GodlessLiu/tpl-cli/config"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tpl-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("V%s\n", config.Version)
	},
}
