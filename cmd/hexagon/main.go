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
	"hexagonal_arch_with_Golang/internal/app/api"
	"hexagonal_arch_with_Golang/internal/app/domain/file"
	"hexagonal_arch_with_Golang/internal/app/domain/notification"
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
	app := api.New(cfg, db, kf, hw, domainFile)

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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c
}

// getLogger set logger config
/*func getLogger(cfg config.Config) logger.Logger {
	var l logger.Log

	l = logger.New(true)

	return l
}*/
