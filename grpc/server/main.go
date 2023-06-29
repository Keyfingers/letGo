package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "letGo/grpc/server/proto" // 替换为你的生成的pb目录路径
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
	// 监听指定端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 注册 Greeter 服务实例到 gRPC 服务器
	pb.RegisterGreeterServer(grpcServer, &server{})

	log.Printf("Server listening at %v", lis.Addr())
	// 启动 gRPC 服务器
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
