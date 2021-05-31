package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"github.com/tetronoz/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Processing a request\n")
	res := req.FirstNumber + req.SecondNumber

	response := &calculatorpb.SumResponse{
		SumResult: res,
	}
	return response, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked\n")

	number := req.GetNumber()
	var k int32 = 2

	for number > 1 {
		if number%k == 0 {
			response := &calculatorpb.PrimeNumberResponse{
				Result: k,
			}
			stream.Send(response)
			number = number / k
		} else {
			k++
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("Starting server ComputeAverage()")

	sum := int32(0)
	count := 0

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			average := float64(sum) / float64(count)
			return stream.SendAndClose(&calculatorpb.AverageResponse{
				Result: average,
			})
		}
		if err != nil {
			log.Fatalf("Error processing ComputeAverage: %v", err)
		}

		sum += msg.GetNumber()
		count++
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("Starting server FindMaximum()")

	current_max := int32(0)

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving from client: %v", err)
			return err
		}
		if msg.Number > current_max {
			current_max = msg.Number
			err = stream.Send(&calculatorpb.MaximumResponse{
				Maximum: current_max,
			})
			if err != nil {
				log.Fatalf("Error sending response to client: %v", err)
				return err
			}
		}
	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Println("Starting server SquareRoot()")
	number := req.GetNumber()

	if number <= 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Calculator")

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
