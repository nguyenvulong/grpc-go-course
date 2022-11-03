package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// id := createBlog(c)
	id := "6342d4caa889ac73d6c7b3d7"
	// readBlog(c, id)
	// readBlog(c, "IdNotExists")
	// updateBlog(c, id)
	listBlog(c)
	// deleteBlog(c, id)
}
