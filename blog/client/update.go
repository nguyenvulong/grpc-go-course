package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("Update a blog")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not clement",
		Title:    "New Title",
		Content:  "New Content",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error while updating blog %v\n", err)
	}

	log.Println("Blog was updated")
}
