package cmd

import (
	notecmd "github.com/Tyler-Arciniaga/DevLog/internal/note"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(noteCommand)
	noteCommand.AddCommand(noteAddCommand)
	noteCommand.AddCommand(noteListCommand)
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
			notecmd.NoteAdd(args[0], args[1])
		} else {
			notecmd.NoteAdd(args[0], "")
		}
	},
}

var noteListCommand = &cobra.Command{
	Use:   "list",
	Short: "view all your previous notes",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		notecmd.NoteList()
	},
}
