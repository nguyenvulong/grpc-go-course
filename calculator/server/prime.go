package main

import (
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Receiving request from %d\n", in.Number)
	// for i := 0; i < 10; i++ {
	// 	stream.Send(&pb.PrimeResponse{
	// 		Decomposition: 33,
	// 	})
	// }

	k := int32(2)
	n := in.Number
	for n > 1 {
		if n%k == 0 {
			stream.Send(&pb.PrimeResponse{Decomposition: k})
			n = n / k
		} else {
			k = k + 1
		}
	}
	return nil
}
