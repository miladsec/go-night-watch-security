package cmd

import (
	"fmt"
	"github.com/miladsec/go-night-watch-security/internal/config"
	"github.com/miladsec/go-night-watch-security/internal/services/migrationService"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}

		migrationService := migrationService.NewMigrationService(cfg)
		migrationService.Run()
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
