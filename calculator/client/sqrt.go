package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) error {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server %v\n", e.Message())

			if e.Code() == codes.InvalidArgument {
				log.Fatal("Probably sent the negative number")
			}
		} else {
			log.Fatalf("non-GRPC error %v\n", err)
		}
	}

	log.Printf("Sqrt %f\n", res.Sqrt)
	return nil
}
