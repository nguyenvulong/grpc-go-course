package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with %v\n", in.Id)
	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	data := &BlogItem{}
	filter := bson.D{{Key: "_id", Value: oid}} //ordered
	res := collection.FindOne(ctx, filter)
	filter2 := bson.M{"_id": oid}
	res2 := collection.FindOne(ctx, filter2) //unordered

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find the blog with the ID provided",
		)
	}
	if err := res2.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find the blog with the ID provided",
		)
	}
	return documentToBlog(data), nil
}
