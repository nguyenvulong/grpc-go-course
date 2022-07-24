package main

import (
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (s *Server) GreatManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreatManyTimesServer) error {
	log.Printf("Greet many times was invoked %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s,  number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}
	return nil
}
