package main

import (
	"io"
	"log"

	pb "github.com/Miranlfk/Go-GRPC/proto"
)

func (s *server) SayHelloBiDiStream(stream pb.GreetService_SayHelloBiDiStreamServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received request with name: %v", request.Name)
		response := &pb.HelloResponse{
			Message: "Hello " + request.Name,
		}
		if err := stream.Send(response); err != nil {
			return err
		}
	}
}