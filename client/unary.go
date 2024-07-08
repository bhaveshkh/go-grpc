package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bhaveshkh/go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	context, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(context, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet server: %v", err)
	}

	log.Printf("%s", res.Message)
}
