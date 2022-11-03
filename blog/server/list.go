package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListLog was invoked")
	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown error %v\n", err),
		)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err = cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decoding data %v\n", err),
			)
		}
		stream.Send(documentToBlog(data))
	}
	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error %v\n", err),
		)
	}
	return nil
}
