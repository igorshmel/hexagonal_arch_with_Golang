package file

import (
	"hexagonal_arch_with_Golang/internal/adapters/ports"
	"hexagonal_arch_with_Golang/pkg/logger"
)

// Endpoint is endpoint for /account routes
type Endpoint struct {
	logger logger.Logger
	app    ports.AppPort
}

// NewEndpoint is constructor for Endpoint
func NewEndpoint(logger logger.Logger, app ports.AppPort) *Endpoint {
	return &Endpoint{logger: logger, app: app}
}
