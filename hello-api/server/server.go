package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sk-develop/grpc-sample/hello-proto"
			
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

// 確認コマンド
// grpc_cli call localhost:9090 hello.Greeter.SayHello 'name: "sk"'
//  grpc_cli ls localhost:9090 hello.Greeter
//  grpc_cli ls localhost:9090 hello.Greeter -l

// 以下を参考にし、統合
// Git:https://github.com/tech-with-moss/go-usermgmt-grpc/blob/main/usermgmt_server/usermgmt_server.go
// Youtube:https://www.youtube.com/watch?v=YudT0nHvkkE
// Git:https://github.com/yassun-youtube/grpc-web-sample/blob/master/api/server.go
// Youtube:「 gRPC Web 」で gRPC 実践！ Go と gRPC で WebAPI を作ってみよう！！