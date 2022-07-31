package main

import (
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	//log.Printf("Receiving request from " )
	sum := 0.0
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// return result
			avg := sum / float64(count)
			stream.SendAndClose(&pb.AvgResponse{Avg: avg})
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving request %v\n", err)
		}
		log.Printf("Received from Client %v\n", req)
		sum = sum + float64(req.Number)
		count = count + 1
	}
}
