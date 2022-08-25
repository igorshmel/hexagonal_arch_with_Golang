package helloworld

import (
	"fmt"

	"hexagonal_arch_with_Golang/pkg/app/api"
	"hexagonal_arch_with_Golang/pkg/config"
)

type HelloWorld struct {
	cfg *config.Config
}

// Check if we actually implement relevant api
var _ api.HelloWorld = &HelloWorld{}

func New(cfg *config.Config) *HelloWorld {
	return &HelloWorld{
		cfg: cfg,
	}
}

func (h *HelloWorld) HelloWorld(greeting, name string) string {
	fmt.Printf("Name is: %s\n", name)
	return fmt.Sprintf("%s %s", greeting, name)
}
