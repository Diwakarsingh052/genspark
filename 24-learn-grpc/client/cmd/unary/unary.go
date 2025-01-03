package main

import (
	pb "client/gen/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	dialOpts := []grpc.DialOption{
		// WithTransportCredentials specifies the transport credentials for the connection
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.NewClient("localhost:5001", dialOpts...)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	user := &pb.User{
		Name:     "test",
		Email:    "d@d.d",
		Password: "<PASSWORD>",
		Roles:    []string{"user"},
	}
	req := &pb.SignupRequest{User: user}
	resp, err := client.Signup(ctx, req)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
