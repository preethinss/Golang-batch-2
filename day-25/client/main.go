package main

import (
	"context"
	"log"

	pb "github.com/PreethiNS/day-25/helloworld"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{Name: "xxx"}

	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}
	log.Printf("Server response: %s", resp.Message)
}
