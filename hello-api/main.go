package main

import (
	"context"
	"log"
	"net"

	// pb "github.com/sk-develop/grpc-sample"
	pb "github.com/sk-develop/grpc-sample/hello-api/hello-proto"
	// ~/go/pkg/mod/
	// github.com/sk-develop/grpc-sample/hello-api@v0.0.0-20210925042332-b8f6ab976a8b/hello-proto

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9090"
)

type HelloworldServer struct {
	pb.UnimplementedGreeterServer
}

func (h HelloworldServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received name: %v", request.Name)
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &HelloworldServer{})
	reflection.Register(server)

	log.Printf("gPRC server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

