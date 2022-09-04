package notificationDomain

import (
	"fmt"

	"hexagonal_arch_with_Golang/pkg/config"
)

type Notification struct {
	cfg     *config.Config
	name    string
	message string
}

// Check if we actually implement relevant api
//var _ api.Notification = &Notification{}

func New(cfg *config.Config) *Notification {
	return &Notification{
		cfg: cfg,
	}
}

func (ths *Notification) Notification(greeting, name string) string {
	fmt.Printf("Name is: %s", name)
	return fmt.Sprintf("%s %s", greeting, name)
}

func (ths *Notification) GetName() string {
	return ths.name
}

func (ths *Notification) GetMassage() string {
	return ths.message
}

func (ths *Notification) SetMassage(message string) {
	ths.message = message
}

func (ths *Notification) SetName(name string) {
	ths.name = name
}
