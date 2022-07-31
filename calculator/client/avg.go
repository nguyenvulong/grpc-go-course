package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) error {
	log.Println("doAvg was invoked")
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error invoking doAvg %v\n", err)
	}

	reqs := []*pb.AvgRequest{
		{Number: 3},
		{Number: 5},
		{Number: 7},
		{Number: 9},
	}

	for _, req := range reqs {
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error sending request %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error closing stream %v\n", err)
	}

	log.Printf("Average %f\n", res.Avg)
	return nil
}
