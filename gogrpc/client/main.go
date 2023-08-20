package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {

	creds := insecure.NewCredentials()

	// Dial creates a client connection to the given target
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	calculatorClient := services.NewCalculatorClient(cc)
	calculatorService := services.NewCalculatorServer(calculatorClient)

	//err = calculatorService.Hello("Napas")
	//err = calculatorService.Fibonacci(5)
	//err = calculatorService.Average(1, 2, 3, 4, 5)
	err = calculatorService.Sum(1, 2, 3, 4, 5)
	if err != nil {
		if grpcErr, ok := status.FromError(err); ok{
			log.Printf("gRPC Error : %v", grpcErr.Message())
			log.Printf("gRPC Error Code : %v", grpcErr.Code())
		}
		log.Fatal(err)
	}

}
