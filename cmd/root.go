package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Initialize function will be called when every command gets called.
func init() {
}

var rootCmd = &cobra.Command{
	Use:   "stockyard",
	Short: "Distributed time-series data server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing.
	},
}

// Execute is the main entry into the application from the command line terminal.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
