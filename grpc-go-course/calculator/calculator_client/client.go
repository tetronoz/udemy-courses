package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tetronoz/grpc-go-course/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error calling Dial(): %v", err)
	}

	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	//doUnary(c)
	//doStreaming(c)
	//doClientStreaming(c)
	//doBiDi(c)
	doErrorUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {

	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Error sending request %v", err)
	}

	fmt.Println(res.SumResult)
}

func doStreaming(c calculatorpb.CalculatorServiceClient) {

	req := &calculatorpb.PrimeNumberRequest{
		Number: 120,
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while calling Recv(): %v", err)
		}
		log.Printf("%v", msg.Result)
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {

	requests := []*calculatorpb.AverageRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error calling ComputeAverage() from client: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving the response: %v", err)
	}
	fmt.Println(response.Result)
}

func doBiDi(c calculatorpb.CalculatorServiceClient) {

	numbers := []int32{1, 5, 3, 6, 2, 20}

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error calling FindMax()")
	}

	waitc := make(chan struct{})
	go func() {
		for _, number := range numbers {
			//fmt.Printf("Sending request: %v\n", req)
			err := stream.Send(&calculatorpb.MaximumRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
			if err != nil {
				log.Fatalf("Error sending client request\n")
			}
		}
		stream.CloseSend()
		close(waitc)
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving response\n")
				break
			}
			fmt.Println(res.Maximum)
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting do SquareRoot Unary RPC")

	number := int32(-10)

	// correct call
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: number})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("Big error")
			return
		}
	}
	fmt.Println(res.NumberRoot)
}
