package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "letGo/grpc/server/proto"
)

func main() {
	// 创建与服务器的连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 创建 Greeter 客户端实例
	client := pb.NewGreeterClient(conn)

	// 构建请求
	request := &pb.HelloRequest{Name: "Alice"}

	// 发送请求
	response, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 输出响应结果
	log.Printf("Greeting: %s", response.Message)
}
