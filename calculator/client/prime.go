package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Printf("doPrime was invoked\n")

	req := &pb.PrimeRequest{
		Number: 120,
	}
	stream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling Prime %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}
		log.Printf("Prime Decomposition %d\n", msg.Decomposition)
	}

}
