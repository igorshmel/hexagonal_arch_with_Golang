package file

import (
	"hexagonal_arch_with_Golang/internal/app"
	"hexagonal_arch_with_Golang/pkg/logger"
)

// Endpoint is endpoint for /account routes
type Endpoint struct {
	logger logger.Logger
	app    app.ApiPort
}

// NewEndpoint is constructor for Endpoint
func NewEndpoint(logger logger.Logger, app app.ApiPort) *Endpoint {
	return &Endpoint{logger: logger, app: app}
}
