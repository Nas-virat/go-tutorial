package services

import (
	context "context"
	"fmt"
	"time"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (s calculatorServer) mustEmbedUnimplementedCalculatorServer() {

}

func (s calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %v at %v", req.Name, req.CreatedAt.AsTime().Local())

	res := HelloResponse{
		Result: result,
	}

	return &res, nil
}

func (s calculatorServer) Fibonacci(req *FibonacciRequest, stream Calculator_FibonacciServer) error {

	fmt.Printf("Service : Fibonacci\n")
	fmt.Printf("Request : %v\n", req.N)
	for n := uint32(0); n < req.N; n++ {
		res := FibonacciResponse{
			Result: fib(n),
		}
		fmt.Printf("Response : %v\n", res.Result)
		stream.Send(&res)
		time.Sleep(time.Second)
	}

	return nil
}

func fib(n uint32) uint32 {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
