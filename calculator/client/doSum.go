package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Num1: 10,
		Num2: 15,
	})

	if err != nil {
		log.Fatalf("Could not sum :%v\n", err)
	}
	log.Printf("Sum result %v\n", res.Sum)

}
