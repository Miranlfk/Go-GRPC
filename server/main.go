package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Miranlfk/Go-GRPC/proto"
)

const (
	port = ":8080"
)

type server struct {
	pb.GreetServiceServer
}


func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})
	log.Printf("Server is running on port: %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil{
		log.Fatalf("Failed to create the server: %v", err)
	}


}