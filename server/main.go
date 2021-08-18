package main

import (
	user_proto "github.com/grpc-http-demo/proto"
	"github.com/grpc-http-demo/server_http"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

const (
	port = "127.0.0.1:50052"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 注册 grpc
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		s := grpc.NewServer()
		user_proto.RegisterAuthServiceServer(s, &User{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 监听 http 请求
	go server_http.HandleHttp()

	wg.Wait()
}
