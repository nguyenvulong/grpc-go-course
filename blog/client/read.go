package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func ReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("ReadBlog was invoked")
	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Printf("Error while reading %v\n", err)
	}
	log.Printf("Log was read %v\n", res)
	return res
}
