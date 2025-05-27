package cmd

import (
	debugcmd "github.com/Tyler-Arciniaga/DevLog/internal/debug"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(debugCommand)
	debugCommand.AddCommand(debugAddCommand)
	debugCommand.AddCommand(debugListCommand)
	debugListCommand.Flags().StringVarP(&Tag, "tag", "t", "", "search for debug task with a specific tag")
	debugCommand.AddCommand(debugSquashCommand)
}

var debugCommand = &cobra.Command{
	Use:   "debug [command]",
	Short: "commands related to your debug log",
}

var debugAddCommand = &cobra.Command{
	Use:   "add [debug note]",
	Short: "add a new debug task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			debugcmd.DebugAdd(args[0], args[1])
		} else {
			debugcmd.DebugAdd(args[0], "")
		}
	},
}

var debugListCommand = &cobra.Command{
	Use:   "list",
	Short: "view all current debug tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		debugcmd.DebugList(Tag)
	},
}

var debugSquashCommand = &cobra.Command{
	Use:   "squash [bug id]",
	Short: "mark a debug task as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		debugcmd.DebugSquash(args[0])
	},
}
