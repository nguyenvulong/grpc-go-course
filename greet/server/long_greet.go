package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("error receiving from client: %v\n", err)
		}
		log.Printf("receiving %v\n", req)
		res += fmt.Sprintf("hello %s\n", req.FirstName)
	}

}
