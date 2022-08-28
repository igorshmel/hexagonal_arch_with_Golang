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
	defaultName = "igor_shmel"
)

var (
	addr = flag.String("addr", "localhost:9000", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
	//url  = "https://scontent-yyz1-1.cdninstagram.com/v/t51.2885-15/302008048_1185966175313439_6089815927638990192_n.jpg?stp=dst-jpg_e35&_nc_ht=scontent-yyz1-1.cdninstagram.com&_nc_cat=103&_nc_ohc=zsyQqv4jdpIAX9Gcjra&edm=ALQROFkBAAAA&ccb=7-5&ig_cache_key=MjkxNDU1NDUwMTI5NDY3ODcwMw%3D%3D.2-ccb7-5&oh=00_AT_a4jiQduB_uAU0Nm6N1KwvYrdWdLgBLB3yf0hXfqDDPw&oe=63130ECE&_nc_sid=30a2ef"
	url = "https://www.meme-arsenal.com/memes/15c768c6dd454a978825bd682340a125.jpg"
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

	/*	rr, err := c.GetGreeting(ctx, &pb2.Input{Name: *name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Rpl from Greeting: %s\n", rr.String())
	*/
	r, err := c.FileForDownload(ctx, &pb.FileReq{Url: url})
	if err != nil {
		log.Fatalf("could not call FileForDownload: %v", err)
	}
	log.Printf("Rpl from FileForDownload: %s\n", r.String())

}
