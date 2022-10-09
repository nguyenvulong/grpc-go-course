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
	// doSum(c)
	// doPrime(c)
	// doAvg(c)
	// doMax(c)
	// doSqrt(c, -10)
	id := createBlog(c)
	ReadBlog(c, id)
	ReadBlog(c, "IdNotExists") 
}
