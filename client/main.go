package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/Miranlfk/Go-GRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	connection, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewGreetServiceClient(connection)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a number (1 to 4) to choose an RPC call, or 0 to exit:")
		fmt.Println("1. SayHello")
		fmt.Println("2. SayHelloServerStream")
		fmt.Println("3. SayHelloClientStream")
		fmt.Println("4. SayHelloBiDiStream")
		fmt.Println("0. EXIT")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}

		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 0 and 4.")
			continue
		}

		names := &pb.ListofNames{
			Names: []string{"Miran", "Luis", "Fernando", "Khalid"},
		}
		switch option {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			callSayHello(client)
		case 2:
			callSayHelloServerStream(client, names)
		case 3:
			callSayHelloClientStream(client, names)
		case 4:
			callSayHelloBiDiStream(client, names)
		default:
			fmt.Println("Invalid option. Please enter a number between 0 and 4.")
		}
	}
}