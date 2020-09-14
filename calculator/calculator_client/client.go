package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shubhamjain2908/grpc-go/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello Go, from client!")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	// after whole main is done executing this (defer) statement will execute
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Printf("Starting to do a Unary RPC for calculator")
	req := &calculatorpb.CalculatorRequest{
		Calc: &calculatorpb.Calculator{
			X: 3,
			Y: 10,
		},
	}
	res, err := c.Calculator(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Calculator RPC: %v", err)
	}
	log.Printf("Response from Calculator: %v", res)
}
