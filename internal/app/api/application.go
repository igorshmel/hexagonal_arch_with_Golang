package api

import (
	"hexagonal_arch_with_Golang/internal/adapters/ports"
	"hexagonal_arch_with_Golang/pkg/config"
)

type Application struct {
	cfg          *config.Config
	db           ports.DbPort
	kf           ports.KafkaPort
	notification Notification
	file         File
}

// Check if we actually implement all the ports.
//var _ ports.AppPort = &Application{}

func New(cfg *config.Config, db ports.DbPort, kf ports.KafkaPort, ntf Notification, file File) *Application {
	return &Application{
		cfg:          cfg,
		db:           db,
		kf:           kf,
		notification: ntf,
		file:         file,
	}
}
