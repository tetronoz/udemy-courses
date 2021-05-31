package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tetronoz/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	//fmt.Printf("Created  client: %f", c)

	//doUnary(c)

	//doServerStreaming(c)
	//doClientStreaming(c)
	//doBiDiStreaming(c)

	doUnaryWithDeadline(c, 5*time.Second)
	doUnaryWithDeadline(c, 1*time.Second)

}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sergey",
			LastName:  "Tolmachev",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sergey",
			LastName:  "Tolmachev",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while calling Recv(): %v", err)
		}

		log.Printf("Response from GreetManyTimes: %v", msg.Result)
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println(("Starting to do Client Streaming"))

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 1",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 2",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 3",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 4",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving LongGreetResponse: %v", err)
	}

	fmt.Println(res)
}

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println(("Starting to do BiDi Client Streaming"))

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 1",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 2",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 3",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Sergey 4",
			},
		},
	}

	// create a stream by invoking client

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error creating a BiDi stream: %v", err)
		return
	}

	waitc := make(chan struct{})

	// send a bunch of messages to the client

	go func() {
		for _, req := range requests {
			//fmt.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive a bunch of messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving BiDi response: %v", err)
				break
			}
			fmt.Println(res.Result)
		}
		close(waitc)
	}()

	//block untill everything is done
	<-waitc
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sergey",
			LastName:  "Tolmachev",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Println("Unexpected error")
			}
		} else {
			log.Fatalf("Big nastry error")
		}
		return
	}

	log.Printf("Response from GreetWithDeadline: %v", res.Result)
}
