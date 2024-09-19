package main

import (
	"context"
	"log"
	pb "training-golang/session-8-introduction-grpc/proto/helloworld/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost: 50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	greeterClient := pb.NewGreeterServiceClient(conn)

	res, err := greeterClient.SayHello(context.Background(), &pb.SayHelloRequest{
		Name: "Joe",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Message)
}
