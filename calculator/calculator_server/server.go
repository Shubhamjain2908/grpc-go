package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shubhamjain2908/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello Go, from server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Greet function was invoked with %v", req)
	X := req.GetCalc().GetX()
	Y := req.GetCalc().GetY()
	result := X + Y
	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}
	return res, nil
}
