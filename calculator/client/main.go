package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	// doSum(c)
	// doPrime(c)
	doAvg(c)
}
