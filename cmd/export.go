package cmd

import (
	"fmt"
	"os"
	"strings"

	exportcmd "github.com/Tyler-Arciniaga/DevLog/internal/export"
	"github.com/spf13/cobra"
)

var sinceDate string

func init() {
	rootCmd.AddCommand(exportCommand)
	exportCommand.Flags().StringVarP(&sinceDate, "since", "s", "", "specified start date of exported logs")
}

var exportCommand = &cobra.Command{
	Use:   "export",
	Short: "export your notes into either markdown or html",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if strings.EqualFold(args[0], "MD") {
			exportcmd.Export("md", sinceDate)
		} else if strings.EqualFold(args[0], "HTML") {
			exportcmd.ExportHTML("html", sinceDate)
		} else {
			fmt.Println("Unrecognized export format, please choose either MD or HTML")
			os.Exit(1)
		}

	},
}
