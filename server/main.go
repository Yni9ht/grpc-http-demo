package main

import (
	"context"
	user_proto "github.com/grpc-http-demo/proto"
	"github.com/grpc-http-demo/server_http"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

const (
	addr = "127.0.0.1:50052"
)

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 注册 grpc
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		s := grpc.NewServer()
		user_proto.RegisterAuthServiceServer(s, &User{})

		// 服务注册名: user-server
		if err := Register(context.TODO(), "user-server", addr); err != nil {
			log.Fatalf("register %s failed:%v", "user-server", err)
		}
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 监听 http 请求
	go server_http.HandleHttp()

	wg.Wait()
}
