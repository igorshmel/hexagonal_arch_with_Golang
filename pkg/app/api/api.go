package api

import (
	"fmt"

	"hexagonal_arch_with_Golang/pkg/adapters/dto/pb"
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

func (ths *Application) SayHello(name string) string {
	err := ths.kf.TestProduce(name, "myTopic")
	if err != nil {
		fmt.Printf("Error  %s\n", err.Error())
	}
	return ths.greeter.HelloWorld(ths.db.GetRandomGreeting(name), name) + "!"
}

func (ths *Application) FileDownload(file *pb.FilePr) error {
	err := ths.db.WriteFileToDownload(file.Name, file.Url)
	if err != nil {
		fmt.Printf("Error DB %s\n", err.Error())
	}
	err = ths.kf.FileProduce(file, "file")
	if err != nil {
		fmt.Printf("Error  %s\n", err.Error())
	}
	return nil
}
