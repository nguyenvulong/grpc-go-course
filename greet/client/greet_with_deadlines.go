package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadlines(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadlines was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	res, err := c.GreetWithDeadlines(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpected GRPC error %v\n", err)

			}
		} else {
			log.Fatalf("non GRPC error %v\n", err)
		}
	}
	log.Printf("Greet with deadlines %v\n", res.Result)

}
