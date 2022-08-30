package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"hexagonal_arch_with_Golang/pkg/adapters/left/grpc"
	kafkaCons "hexagonal_arch_with_Golang/pkg/adapters/left/kafka"
	pgql "hexagonal_arch_with_Golang/pkg/adapters/right/db"
	"hexagonal_arch_with_Golang/pkg/adapters/right/kafka"
	"hexagonal_arch_with_Golang/pkg/app/application"
	fileDomain "hexagonal_arch_with_Golang/pkg/app/domain/file"
	"hexagonal_arch_with_Golang/pkg/dto/pb"

	"hexagonal_arch_with_Golang/pkg/app/domain/notification"
	"hexagonal_arch_with_Golang/pkg/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.New(ctx)
	if err := cfg.Read(); err != nil {
		if errors.Is(err, config.ErrExit) {
			return
		}
		panic(err)
	}

	// Create Right adapters
	//db := memory.New(cfg)

	db, err := pgql.New(cfg, true)
	if err != nil {
		fmt.Printf("DB connect error: %s\n", err)
	}

	kf, err := kafka.New(cfg)
	if err != nil {
		fmt.Printf("DB connect error: %s\n", err)
	}

	// Create the Domain layer and Application layer
	hw := notificationDomain.New(cfg)
	domainFile := fileDomain.New(cfg)
	app := application.New(cfg, db, kf, hw, domainFile)

	// Create left adapters
	rpc, err := grpc.New(cfg, app)
	if err != nil {
		fmt.Println("Could not create GRPC: ", err)
	} else {
		go rpc.Run()
	}

	kafkaC, err := kafkaCons.New(cfg, "667", app)
	if err != nil {
		fmt.Println("Could not create kafkaCons: ", err)
	} else {
		kafkaC.Consumer([]string{"file", "notification"}, (&pb.FileProducer{}).ProtoReflect().Type(), (&pb.NotificationProducer{}).ProtoReflect().Type())
	}

	/*
		nc, err := nats.New(cfg, app)
		if err != nil {
			fmt.Println("Could not create NATS: ", err)
		} else {
			go nc.Run()
		}
	*/
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c
}
