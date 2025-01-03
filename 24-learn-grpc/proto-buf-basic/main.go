package main

import (
	"fmt"
	pb "proto-buf-basic/proto"
)

func main() {
	b := pb.BlogRequest{
		BlogId:  1,
		Name:    "some random name",
		Title:   "random title",
		Content: "random content",
	}
	fmt.Println(b)
	fmt.Println(b.String())
	fmt.Println(b.GetName())
	//proto.Marshal(b)
}
