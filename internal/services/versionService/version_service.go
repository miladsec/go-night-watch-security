package versionService

import (
	"fmt"
	"github.com/miladsec/go-night-watch-security/internal/config"
)

type VersionService struct {
	config *config.Config
}

func NewVersionService(cfg *config.Config) *VersionService {
	return &VersionService{
		config: cfg,
	}
}

func (s *VersionService) Run() {
	fmt.Printf("you are using %s gnws version.", s.config.Version)
}
