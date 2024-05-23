package main

import (
	"context"
	pb "github.com/Miranlfk/Go-GRPC/proto"
)

func (s *server) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	
	return &pb.HelloResponse{
		Message: "Hello " ,
		}, nil
}