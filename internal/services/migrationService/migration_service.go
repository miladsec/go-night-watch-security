package migrationService

import (
	"fmt"
	"github.com/miladsec/go-night-watch-security/internal/config"
)

type MigrationService struct {
	config *config.Config
}

func NewMigrationService(cfg *config.Config) *MigrationService {
	return &MigrationService{
		config: cfg,
	}
}

func (s *MigrationService) Run() {
	fmt.Printf("Migration is running now...")
}
