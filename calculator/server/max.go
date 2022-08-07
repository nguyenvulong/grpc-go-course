package main

import (
	"io"
	"log"
	"math"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max request is invoked")
	maxnum := int32(math.MinInt32)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving request from client %v\n", err)
		}

		log.Println(req.Number)
		if maxnum < req.Number {
			maxnum = req.Number
		}
		err = stream.Send(&pb.MaxResponse{
			Maxnum: maxnum,
		})
		if err != nil {
			log.Fatalf("Error sending max number to client %v\n", err)
		}
	}
	return nil
}
