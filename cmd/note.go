package cmd

import (
	notecmd "github.com/Tyler-Arciniaga/DevLog/internal/note"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(noteCommand)
	noteCommand.AddCommand(noteAddCommand)
}

var noteCommand = &cobra.Command{
	Use:   "note [command]",
	Short: "Command related to devlog notes",
}

var noteAddCommand = &cobra.Command{
	Use:   "add [your note here]",
	Short: "Used to add a new note to your log file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			notecmd.Note(args[0], args[1])
		} else {
			notecmd.Note(args[0], "")
		}
	},
}
