package memory

import (
	"math/rand"

	"hexagonal_arch_with_Golang/pkg/config"
)

type Adapter struct {
	cfg       *config.Config
	greetings []string
}

//var _ ports.DbPort = &Adapter{}

func New(cfg *config.Config) *Adapter {
	return &Adapter{
		cfg:       cfg,
		greetings: []string{"Hello,", "Hi,", "Ahoy,", "Aloha,"},
	}
}

func (ths *Adapter) GetRandomNotification() string {
	return ths.greetings[rand.Intn(len(ths.greetings))]
}
