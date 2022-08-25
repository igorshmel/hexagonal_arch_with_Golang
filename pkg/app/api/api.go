package api

import (
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/ports"
)

type Application struct {
	cfg     *config.Config
	db      ports.DbPort
	kf      ports.KafkaPort
	greeter HelloWorld
}

// Check if we actually implement all the ports.
var _ ports.APIPort = &Application{}

func New(cfg *config.Config, db ports.DbPort, kf ports.KafkaPort, greeter HelloWorld) *Application {
	return &Application{
		cfg:     cfg,
		greeter: greeter,
		db:      db,
		kf:      kf,
	}
}

func (a *Application) SayHello(name string) string {
	a.kf.TestProduce(name, "myTopic")
	return a.greeter.HelloWorld(a.db.GetRandomGreeting(name), name) + "!"
}
