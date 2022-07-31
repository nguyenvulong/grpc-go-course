package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	fmt.Println("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "ABC"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)

	}

	for _, req := range reqs {
		log.Printf("Sending request %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet %v\n", err)
	}

	log.Printf("Long Greet %s\n", res.Result)
}
