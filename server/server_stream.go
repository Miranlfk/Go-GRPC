package main

import (
	"log"
	"time"

	pb "github.com/Miranlfk/Go-GRPC/proto"
)	

func SayHelloServerStream(req *pb.ListofNames, stream pb.GreetService_SayHelloServerStreamServer) error {
	log.Printf("Recieved request with names: %v", req.Names)

	for _, name := range req.Names {
		if err := stream.Context().Err(); err != nil {
			log.Printf("Context error: %v", err)
			return err
		}

		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		
		time.Sleep(2 * time.Second)
	}
	return nil
}