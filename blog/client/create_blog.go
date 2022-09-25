package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("Create Blog was invoked")

	blog := &pb.Blog{
		AuthorId: "Clement",
		Title:    "My Blog",
		Content:  "Content of the blog",
	}
	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected Error %v\n", err)
	}

	log.Printf("Blog has been created %s\n", res.Id)
	return res.Id
}
