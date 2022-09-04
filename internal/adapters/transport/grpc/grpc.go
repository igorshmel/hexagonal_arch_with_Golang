package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"hexagonal_arch_with_Golang/internal/adapters/ports"
	"hexagonal_arch_with_Golang/pkg/config"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

type Adapter struct {
	cfg    *config.Config
	app    ports.AppPort
	listen net.Listener
}

func New(cfg *config.Config, app ports.AppPort) (*Adapter, error) {
	ret := &Adapter{cfg: cfg, app: app}

	listen, err := net.Listen(cfg.GRPC.Network, cfg.GRPC.Address)
	if err != nil {
		return nil, err
	}
	ret.listen = listen

	return ret, nil
}

func (ths *Adapter) Run() {
	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, ths)
	if err := grpcServer.Serve(ths.listen); err != nil {
		ths.cfg.Logger.Error("failed to serve gRPC server over address %s: %v", ths.cfg.GRPC.Address, err)
	}
}

func (ths *Adapter) File(ctx context.Context, fileReq *pb.FileReq) (*pb.FileRpl, error) {
	if fileReq == nil {
		return nil, fmt.Errorf("input is mandatory")
	}

	err := ths.app.NewFile(fileReq.Url)
	if err != nil {
		return nil, fmt.Errorf("filed Call to app.NewFile()")
	}

	return &pb.FileRpl{}, nil
}
