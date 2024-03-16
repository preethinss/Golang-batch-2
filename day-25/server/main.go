package main

import (
	"context"
	"log"
	"net"

	pb "github.com/PreethiNS/day-25/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type server struct{
	pb.UnimplementedGreeterServer
}

func(s *server)SayHello(ctx context.Context,in *pb.HelloRequest)(*pb.HelloReply,error){
	grpclog.Infof("Recieved request with name :%s\n",in.Name)
	return &pb.HelloReply{Message : "hello "+ in.Name},nil
}

func main() {
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})

	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	log.Println("server listening on port :50051")

	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to servr : %v", err)
	}

}
