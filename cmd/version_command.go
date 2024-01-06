package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of your CLI tool",
	Long:  "Print the version number and additional information about your CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Your custom command logic here
		fmt.Println("you are using 0.0.0 gnws version.")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
