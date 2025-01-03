package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "server/gen/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":5001")

	if err != nil {
		panic(err)
	}

	//NewServer creates a gRPC server which has no service registered
	// creating an instance of the server
	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &userService{})

	//exposing gRPC service to be tested by postman
	reflection.Register(s)

	err = s.Serve(listener) // run the gRPC server
	if err != nil {
		panic(err)
	}
}
