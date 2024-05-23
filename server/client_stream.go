package main

import (
	"io"
	"log"

	pb "github.com/Miranlfk/Go-GRPC/proto"
)

func (s *server) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error{
	var messages []string
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageResponse{Messages: messages})
			
		 }
		if err != nil {
			return err
		}
		log.Printf("Received request with name: %v", request.Name)
		messages = append(messages, "Hello", request.Name)
	}
}