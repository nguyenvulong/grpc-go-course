package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("Greet many times client was invoked")
	req := &pb.GreetRequest{
		FirstName: "Clement",
	}

	stream, err := c.GreatManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error calling Greet Many Times %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream %v\n", err)
		}
		log.Printf("Greet many times %s\n", msg.Result)
	}
}
