package main

import (
	"log"

	pb "./hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Alex",
	})

	if err != nil {
		log.Fatalf("could not hello: %v", err)
	}
	log.Println(r.Message)
}
