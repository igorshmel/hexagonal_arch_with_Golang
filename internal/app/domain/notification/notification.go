package notificationDomain

import (
	"fmt"

	"hexagonal_arch_with_Golang/pkg/config"
)

type Notification struct {
	cfg *config.Config
}

// Check if we actually implement relevant api
//var _ api.Notification = &Notification{}

func New(cfg *config.Config) *Notification {
	return &Notification{
		cfg: cfg,
	}
}

func (h *Notification) Notification(greeting, name string) string {
	fmt.Printf("Name is: %s\n", name)
	return fmt.Sprintf("%s %s", greeting, name)
}
