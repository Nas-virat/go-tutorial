package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CalculatorService interface {
	Hello(name string) error
	Fibonacci(n uint32) error
}

type calculatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorServer(calculatorClient CalculatorClient) CalculatorService {
	return calculatorService{calculatorClient}
}

func (base calculatorService) Hello(name string) error {
	req := HelloRequest{
		Name:      name,
		CreatedAt: timestamppb.Now(),
	}

	res, err := base.calculatorClient.Hello(context.Background(), &req)
	if err != nil {
		return err
	}

	fmt.Print("Service : Hello\n")
	fmt.Printf("Request : %v\n", req.Name)
	fmt.Printf("Response : %v\n", res.Result)

	return nil
}

func (base calculatorService) Fibonacci(n uint32) error {
	req := FibonacciRequest{
		N: n,
	}

	// Declare a timeout of the service
	// If the stream take to long to response, it will be cancel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//if timout is reached, cancel will be called
	defer cancel()

	stream, err := base.calculatorClient.Fibonacci(ctx, &req)
	if err != nil {
		return err
	}

	fmt.Printf("Service : Fibonacci\n")
	fmt.Printf("Request : %v\n", req.N)
	for {
		res, err := stream.Recv()

		// if the stream is finish( End of File )
		// break the loop
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Response : %v\n", res.Result)
	}

	return nil
}
