package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	ppb "hexagonal_arch_with_Golang/pkg/adapters/dto/pb"
	"hexagonal_arch_with_Golang/pkg/adapters/left/grpc"
	kafkaCons "hexagonal_arch_with_Golang/pkg/adapters/left/kafka"
	pgql "hexagonal_arch_with_Golang/pkg/adapters/right/db"
	"hexagonal_arch_with_Golang/pkg/adapters/right/kafka"
	//"hexagonal_arch_with_Golang/pkg/adapters/right/memory"
	"hexagonal_arch_with_Golang/pkg/app/api"
	"hexagonal_arch_with_Golang/pkg/app/core/helloworld"
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
	hw := helloworld.New(cfg)
	app := api.New(cfg, db, kf, hw)

	// Create left adapters
	rpc, err := grpc.New(cfg, app)
	if err != nil {
		fmt.Println("Could not create GRPC: ", err)
	} else {
		go rpc.Run()
	}

	kafkaC, err := kafkaCons.New(cfg, "667")
	if err != nil {
		fmt.Println("Could not create kafkaCons: ", err)
	} else {
		go kafkaC.TestConsumer([]string{"file", "myTopic"}, (&ppb.FilePr{}).ProtoReflect().Type())
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
