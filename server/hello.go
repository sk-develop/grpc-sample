package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/sk-develop/grpc-sample/hello-proto"
			
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type HelloworldServer struct {
	pb.UnimplementedGreeterServer
}

func (h HelloworldServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Helloooo " + request.Name}, nil
}

func main() {
	port := ":9090"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &HelloworldServer{})
	reflection.Register(server)

	log.Printf("start gPRC server")
	server.Serve(lis)
}

// func (s *GreeterServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	var user_id int32 = int32(rand.Intn(100))
// 	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
// }

// func main() {
// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterGreeterServer(s, &GreeterServer{})
// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }