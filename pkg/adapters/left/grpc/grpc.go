package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
	"hexagonal_arch_with_Golang/pkg/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
	cfg    *config.Config
	app    ports.AppPort
	listen net.Listener
}

func New(cfg *config.Config, app ports.AppPort) (*Adapter, error) {
	ret := &Adapter{app: app}

	listen, err := net.Listen(cfg.GRPC.Network, cfg.GRPC.Address)
	if err != nil {
		return nil, err
	}
	ret.listen = listen

	return ret, nil
}

func (a *Adapter) Run() {
	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, a)
	if err := grpcServer.Serve(a.listen); err != nil {
		log.Fatalf("failed to serve gRPC server over address %s: %v", a.cfg.GRPC.Address, err)
	}
}

func (a *Adapter) File(ctx context.Context, fileReq *pb.FileReq) (*pb.FileRpl, error) {
	if fileReq == nil {
		return nil, fmt.Errorf("input is mandatory")
	}

	err := a.app.AppFile(fileReq.Url)
	if err != nil {
		return nil, fmt.Errorf("filed Call to File()")
	}

	return &pb.FileRpl{}, nil
}
