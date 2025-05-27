package cmd

import (
	"fmt"
	"os"
	"strings"

	exportcmd "github.com/Tyler-Arciniaga/DevLog/internal/export"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exportCommand)
}

var exportCommand = &cobra.Command{
	Use:   "export",
	Short: "export your notes into either markdown or html",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if strings.EqualFold(args[0], "MD") {
			exportcmd.Export("md", "")
		} else if strings.EqualFold(args[0], "HTML") {
			exportcmd.ExportHTML("html", "")
		} else {
			fmt.Println("Unrecognized export format, please choose either MD or HTML")
			os.Exit(1)
		}

	},
}
