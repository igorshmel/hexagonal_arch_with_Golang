package api

import (
	"fmt"

	common "hexagonal_arch_with_Golang/pkg/adapters/common/file"
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/ports"
)

type Application struct {
	cfg        *config.Config
	db         ports.DbPort
	kf         ports.KafkaPort
	helloWorld HelloWorld
	post       Post
}

// Check if we actually implement all the ports.
var _ ports.APIPort = &Application{}

func New(cfg *config.Config, db ports.DbPort, kf ports.KafkaPort, helloWorld HelloWorld, post Post) *Application {
	return &Application{
		cfg:        cfg,
		db:         db,
		kf:         kf,
		helloWorld: helloWorld,
		post:       post,
	}
}

func (ths *Application) SayHello(name string) string {
	err := ths.kf.TestProduce(name, "myTopic")
	if err != nil {
		fmt.Printf("Error  %s\n", err.Error())
	}
	return ths.helloWorld.HelloWorld(ths.db.GetRandomGreeting(name), name) + "!"
}

func (ths *Application) FileDownload(fileUrl string) error {
	ths.post.SetFileUrl(fileUrl)
	ths.post.FileParseUrl()

	err := ths.db.WriteFileToDownload(ths.post.GetFileName(), ths.post.GetFileUrl(), ths.post.GetFileStatus())
	if err != nil {
		fmt.Printf("Error DB %s\n", err.Error())
	}

	err = ths.kf.FileProduce(ths.post.GetFileName(), ths.post.GetFileUrl(), ths.post.GetFileStatus(), "file")
	if err != nil {
		fmt.Printf("Error  %s\n", err.Error())
	}
	return nil
}

func (ths *Application) Download(fileUrl, fileName string) {

	fmt.Printf("\ngorutina for Download %s\n", fileName)

	err := common.DownloadFile("instagram.jpg", fileUrl)
	if err != nil {
		fmt.Printf("Error  %s\n", err.Error())
	} else {
		fmt.Printf("\nфайл удалось скачать %s\n", fileName)
	}
}
