package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"log"
	"time"

	user_proto "github.com/grpc-http-demo/proto"
	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50052"
	defaultName = "world"
)

func main() {
	cli, err := clientv3.NewFromURL("http://localhost:2379")
	if err != nil {
		panic(err)
	}
	builder, err := resolver.NewBuilder(cli)
	if err != nil {
		panic(err)
	}
	// 获取服务
	conn, err := grpc.DialContext(context.TODO(), "etcd:///service/user-server",
		grpc.WithResolvers(builder),
		grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	c := user_proto.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	r, err := c.UserInfo(ctx, &user_proto.UserInfoReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("response: %v \n", r)
	defer cancel()
}
