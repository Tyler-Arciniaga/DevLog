package cmd

import (
	initcmd "github.com/Tyler-Arciniaga/DevLog/internal/init"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize devlog directory",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		initcmd.Init()
	},
}
