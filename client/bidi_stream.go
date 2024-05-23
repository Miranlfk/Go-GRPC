package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Miranlfk/Go-GRPC/proto"
)

func callSayHelloBiDiStream(client pb.GreetServiceClient, names *pb.ListofNames) {
	log.Printf("Bidirectional Streaming has started")
	stream, err := client.SayHelloBiDiStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send the names: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for{
			messages, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while Streaming: %v", err)
			}
			log.Println(messages)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		request := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(request); err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming has ended")
}