package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while strreaming %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Clement"},
		{FirstName: "Marie"},
		{FirstName: "Test"},
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			err = stream.Send(req)
			if err != nil {
				log.Fatalf("Error sending request %v\n", err)
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
				//log.Fatalf("Error receving stream %v\n", err)
				log.Printf("Error receving stream %v\n", err)
				break

			}
			log.Printf("Received %v", res)
		}
		close(waitc)
	}()
	<-waitc

}
