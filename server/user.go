package main

import (
	"context"
	"fmt"

	user_proto "github.com/grpc-http-demo/proto"
)

type User struct {
	user_proto.UnimplementedAuthServiceServer
}

func (u *User) UserInfo(ctx context.Context, req *user_proto.UserInfoReq) (*user_proto.UserInfoRes, error) {
	fmt.Printf("hahhahhhhahhahahahaahhah")
	fmt.Println("================")
	return &user_proto.UserInfoRes{}, nil
}
