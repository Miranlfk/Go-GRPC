package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Miranlfk/Go-GRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)

	}
	log.Printf("Response: %s", response.Message)
	
}