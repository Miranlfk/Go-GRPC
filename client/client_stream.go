package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Miranlfk/Go-GRPC/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.ListofNames){
	log.Println("Client Streaming has started")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send the names: %v", err)
	}

	for _, name := range names.Names {
		request := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("Could not send the request: %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	response, err := stream.CloseAndRecv()
	log.Printf("Client Streaming has ended")
	if err != nil {
		log.Fatalf("Could not receive the response: %v", err)
	}
	log.Printf("Response: %v", response.Messages)
}