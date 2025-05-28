package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "See current version of DevLog installed on local machine",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DevLog 1.0.0")
	},
}
