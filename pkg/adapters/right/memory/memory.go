package memory

import (
	"math/rand"

	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/ports"
)

type Adapter struct {
	cfg       *config.Config
	greetings []string
}

var _ ports.DbPort = &Adapter{}

func New(cfg *config.Config) *Adapter {
	return &Adapter{
		cfg:       cfg,
		greetings: []string{"Hello,", "Hi,", "Ahoy,", "Aloha,"},
	}
}

func (ths *Adapter) GetRandomGreeting(name string) string {
	return ths.greetings[rand.Intn(len(ths.greetings))]
}

func (ths *Adapter) GetGreetings() []string {
	return ths.greetings
}
func (ths *Adapter) WriteFileToDownload(fileName, fileUrl, fileStatus string) error {
	return nil
}
