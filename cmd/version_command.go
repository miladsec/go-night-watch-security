package cmd

import (
	"fmt"
	"github.com/miladsec/go-night-watch-security/internal/config"
	"github.com/miladsec/go-night-watch-security/internal/service/versionService"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of your CLI tool",
	Long:  "Print the version number and additional information about your CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		versionService := versionService.NewVersionService(cfg)
		versionService.Run()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
