package main

import (
	"context"
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
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user_proto.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	r, err := c.UserInfo(ctx, &user_proto.UserInfoReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("response: %v \n", r)
	defer cancel()
}
