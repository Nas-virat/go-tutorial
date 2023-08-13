package services

import (
	context "context"
	"fmt"
)

type calculatorService struct {
}

func NewCalculatorServer() CalculatorServer {
	return calculatorService{}
}

func (s calculatorService) mustEmbedUnimplementedCalculatorServer() {

}

func (s calculatorService) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	result := fmt.Sprintf("Hello %v", req.Name)
	res := HelloResponse{
		Result: result,
	}

	return &res, nil
}
