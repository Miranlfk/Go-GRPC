package main 

import(
	"context"
	"log"
	"io"
	"time"

	pb "github.com/Miranlfk/Go-GRPC/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.ListofNames) {
	log.Printf("Server Streaming has started")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.SayHelloServerStream(ctx, names)
	if err != nil {
		log.Fatalf("Failed to call SayHelloServerStream: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive the response: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming has finished")
}

