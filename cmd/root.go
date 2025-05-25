/*
Copyright © 2025 Tyler Arciniaga tyarciniaga@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devlog",
	Short: "A lightweight CLI tool to track development notes and debugging progress",
	Long: `DevLog CLI is a lightweight tool for managing project notes, debug checklists, 
and development logs—all from the terminal.

Use it to:
- Add timestamped notes about what you've worked on
- Track and check off debug tasks
- Initialize a .devlog folder for storing logs in JSON or YAML
- (Optional) Summarize Git commits to review your progress

Everything is stored locally, version-controllable, 
and designed to fit smoothly into your development workflow.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.DevLog.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
