package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Printf("doMax was invoked\n")

	reqs := []*pb.MaxRequest{
		{Number: 5},
		{Number: 9},
		{Number: 8},
		{Number: 12},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error calling Max %v\n", err)
	}
	awaitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			err = stream.Send(req)
			if err != nil {
				log.Fatalf("Error sending number to server %v\n", err)
			}
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving result from server %v\n", err)
			}
			log.Printf("Max Number is %v\n", res)
		}
		close(awaitc)
	}()
	<-awaitc

}
