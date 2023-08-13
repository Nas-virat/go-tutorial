package main

import (
	"client/services"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	err = calculatorService.Hello("Napas")
	if err != nil {
		log.Fatal(err)
	}

}
