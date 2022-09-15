package api

import (
	"hexagonal_arch_with_Golang/internal/adapters/ports"
	"hexagonal_arch_with_Golang/internal/domain"
	"hexagonal_arch_with_Golang/pkg/config"
)

type Service struct {
	cfg  *config.Config
	db   ports.DbPort
	file domain.File
}

// Check if we actually implement all the ports.
//var _ ports.AppPort = &Service{}

func New(cfg *config.Config, db ports.DbPort, file domain.File) *Service {
	return &Service{
		cfg:  cfg,
		db:   db,
		file: file,
	}
}
