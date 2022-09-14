package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"hexagonal_arch_with_Golang/internal/adapters/stage/db"
	"hexagonal_arch_with_Golang/internal/adapters/stage/kafka"
	"hexagonal_arch_with_Golang/internal/adapters/transport/grpc"
	kafkaCons "hexagonal_arch_with_Golang/internal/adapters/transport/kafka"
	"hexagonal_arch_with_Golang/internal/adapters/transport/rest"
	"hexagonal_arch_with_Golang/internal/domain/file"
	"hexagonal_arch_with_Golang/internal/domain/notification"
	"hexagonal_arch_with_Golang/internal/service/api"
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
	"hexagonal_arch_with_Golang/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New(true)
	cfg := config.New(ctx)
	if err := cfg.Read(); err != nil {
		if errors.Is(err, config.ErrExit) {
			return
		}
		panic(err)
	}
	cfg.Logger = log

	// Create Right adapters
	//mem := memory.New(cfg)

	db, err := pgql.New(cfg, true)
	if err != nil {
		log.Error("DB connect error: %s", err)
	}

	kf, err := kafka.New(cfg)
	if err != nil {
		fmt.Printf("Kafka init error: %s", err)
	}

	// Create the Domain layer and Service layer
	hw := notificationDomain.New(cfg)
	domainFile := fileDomain.New(cfg)
	app := api.New(cfg, db, kf, hw, domainFile)

	// Create left adapters
	rpcAdapter, err := grpc.New(cfg, app)
	if err != nil {
		fmt.Println("Could not create GRPC: ", err)
	} else {
		go rpcAdapter.Run()
	}

	// Create left adapters
	restAdapter, err := rest.New(cfg, app)
	if err != nil {
		fmt.Println("Could not start rest Engine: ", err)
	} else {
		go restAdapter.Run()
		restAdapter.FileHandlers()
	}

	kafkaAdapter, err := kafkaCons.New(cfg, "667", app)
	if err != nil {
		fmt.Println("Could not create kafkaCons: ", err)
	} else {
		go kafkaAdapter.Consumer(
			[]string{"file", "notification"},
			(&pb.FileProducer{}).ProtoReflect().Type(),
			(&pb.NotificationProducer{}).ProtoReflect().Type(),
		)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c

}
