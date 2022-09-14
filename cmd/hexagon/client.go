package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hexagonal_arch_with_Golang/pkg/dto/pb"
)

const (
	defaultName = "default_name"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	url  = "https://thispersondoesnotexist.com/image"
	//url  = "https://www.meme-arsenal.com/memes/15c768c6dd454a978825bd682340a125.jpg"
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
	c := pb.NewFileServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.File(ctx, &pb.FileReq{Url: url})
	if err != nil {
		log.Fatalf("could not call File: %v", err)
	}
	log.Printf("Rpl from File: %s", r.String())

}
