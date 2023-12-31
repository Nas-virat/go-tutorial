package services

import (
	context "context"
	"fmt"
	"io"
	"time"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type calculatorServer struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorServer{}
}

func (s calculatorServer) mustEmbedUnimplementedCalculatorServer() {

}

func (s calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	if req.Name == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"name is required",
		)
	}

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

func (s calculatorServer) Average(stream Calculator_AverageServer) error {
	sum := 0.0
	count := 0.0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}

	res := AverageResponse{
		Result: sum / count,
	}

	return stream.SendAndClose(&res)
}

func (s calculatorServer) Sum(stream Calculator_SumServer) error {
	sum := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		sum += req.Number
		res := SumResponse{
			Result: sum,
		}
		err = stream.Send(&res)
		if err != nil {
			return err
		}
	}

	return nil
}

func fib(n uint32) uint32 {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
