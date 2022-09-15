package rest

import (
	"github.com/gin-gonic/gin"
	"hexagonal_arch_with_Golang/internal/adapters/transport/rest/file"
	"hexagonal_arch_with_Golang/internal/service"
	"hexagonal_arch_with_Golang/pkg/config"
)

type Adapter struct {
	cfg     *config.Config
	service service.ApiPort
	router  *gin.Engine
}

func New(cfg *config.Config, app service.ApiPort) (*Adapter, error) {
	router := gin.New()
	ret := &Adapter{cfg: cfg, service: app, router: router}

	return ret, nil
}

func (ths *Adapter) Run() {
	err := ths.router.Run(ths.cfg.Env.ListenAddr)
	if err != nil {
		ths.cfg.Logger.Error("GIN run failed, error: %s", err)
	}
}

func (ths *Adapter) FileHandlers() {

	// set route handlers
	apiGroup := ths.router.Group("/api")

	// account routes
	accountEndpoint := file.NewEndpoint(ths.cfg.Logger, ths.service)
	apiGroup.POST("/file_url/", accountEndpoint.NewFileHandler())

	ths.cfg.Logger.Info("route File handlers are installed...")
}
