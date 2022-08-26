package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hexagonal_arch_with_Golang/pkg/adapters/left/grpc/pb"
)

const (
	defaultName = "igor_shmel"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	url  = "http://fileurl.com"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := pb.NewHelloWorldClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	/*
		r, err := c.GetGreeting(ctx, &pb.Input{Name: *name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	*/
	r, err := c.FileForDownload(ctx, &pb.FileReq{Url: url})
	if err != nil {
		log.Fatalf("could not call FileForDownload: %v", err)
	}
	log.Printf("Rpl from FileForDownload: %s\n", r.String())
}
