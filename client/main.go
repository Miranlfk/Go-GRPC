package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/Miranlfk/Go-GRPC/proto"
)

const (
	port = ":8080"
)

func main() {
	connection, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect : %v", err)
	}

	defer connection.Close()

	client := pb.NewGreetServiceClient(connection)

	names := &pb.ListofNames{
		Names: []string{"Miran", "Luis", "Fernando", "Khalid"},
	}
	// callSayHello(client)
	callSayHelloServerStream(client, names)
	
}