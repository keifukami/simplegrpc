package main

import (
	"context"
	"fmt"
	pb "github.com/keifukami/simplegrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"time"
)

const (
	serverAddress = "localhost:50051"
)

func main() {

	var err error

	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(serverAddress, opts...)
	if err != nil {
		fmt.Printf("failed to dial: %s\n", serverAddress)
		panic(err)
	}
	defer (func(c *grpc.ClientConn) {
		if err := c.Close(); err != nil {
			fmt.Printf("failed to close connection: %#v\n", err)
		}
	})(conn)

	echoClient := pb.NewEchoClient(conn)
	calcClient := pb.NewCalculatorClient(conn)

	fmt.Println("-----")
	fmt.Println("call proto.Echo/OneEcho")
	err = callOneEcho(echoClient, "myname", "yourname", "Hello, World!")
	if err != nil {
		fmt.Printf("OneEcho failed: %#v\n", err)
	}

	fmt.Println("-----")
	fmt.Println("call proto.Echo/MultiEcho")
	err = callMultiEcho(echoClient, "myname", "yourname", "Hello, World!", 4, 2)
	if err != nil {
		fmt.Printf("MultiEcho failed: %#v\n", err)
	}

	fmt.Println("-----")
	fmt.Println("call proto.Calculator/Add")
	err = callAdd(calcClient, []int32{10, 5, 200})
	if err != nil {
		fmt.Printf("Add failed: %#v\n", err)
	}

	fmt.Println("-----")
	fmt.Println("call proto.Calculator/AddInteractive")
	err = callAddInteractive(calcClient, []int32{10, 5, 200})
	if err != nil {
		fmt.Printf("AddInteractive failed: %#v\n", err)
	}

}

func callOneEcho(client pb.EchoClient, myName string, yourName string, message string) error {

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	msg := &pb.Message{
		SourceName:      myName,
		DestinationName: yourName,
		Body:            message,
	}

	var resp *pb.Message
	resp, err = client.OneEcho(ctx, msg)
	if err != nil {
		return err
	}
	fmt.Printf("receive echo: %s\n", resp.String())

	return nil

}

func callMultiEcho(
	client pb.EchoClient,
	myName string,
	yourName string,
	message string,
	repeats uint32,
	delayInSec uint32,
) error {

	var err error

	expectedDuration := (time.Duration)(delayInSec * (repeats - 1))
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second+expectedDuration)
	defer cancel()

	req := &pb.MultiEchoRequest{
		Message: &pb.Message{
			SourceName:      myName,
			DestinationName: yourName,
			Body:            message,
		},
		Repeats:        &repeats,
		DelayInSeconds: &delayInSec,
	}

	var stream pb.Echo_MultiEchoClient
	stream, err = client.MultiEcho(ctx, req)
	if err != nil {
		return err
	}

	var msg *pb.Message
	for {

		msg, err = stream.Recv()

		if err == io.EOF {
			fmt.Println("echo stream closed")
			return nil
		}

		if err != nil {
			fmt.Printf("receive error: %#v", err)
			return err
		}

		fmt.Printf("receive echo: %s\n", msg.String())

	}

}

func callAdd(client pb.CalculatorClient, operands []int32) error {

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var addClient pb.Calculator_AddClient
	addClient, err = client.Add(ctx)
	if err != nil {
		fmt.Printf("failed to open session for Add: %#v\n", err)
		return err
	}

	for _, operand := range operands {
		val := &pb.Value{
			Number: operand,
		}
		err = addClient.Send(val)
		if err != nil {
			fmt.Printf("failed to send: %#v\n", err)
			return err
		}
	}

	var sumValue *pb.Value
	sumValue, err = addClient.CloseAndRecv()
	if err != nil {
		fmt.Printf("failed to receive result: %#v\n", err)
		return err
	}

	fmt.Printf("get sum: %d\n", sumValue.Number)

	return nil

}

func callAddInteractive(client pb.CalculatorClient, operands []int32) error {

	var err error

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var addClient pb.Calculator_AddInteractiveClient
	addClient, err = client.AddInteractive(ctx)
	if err != nil {
		fmt.Printf("failed to open session for Add: %#v\n", err)
		return err
	}

	var lastResult int32 = 0
	for _, operand := range operands {
		val := &pb.Value{
			Number: operand,
		}
		err = addClient.Send(val)
		if err != nil {
			fmt.Printf("failed to send: %#v\n", err)
			return err
		}

		var currentSum *pb.Value
		currentSum, err = addClient.Recv()
		if err != nil {
			fmt.Printf("failed to receive intermediate result: %#v\n", err)
			return err
		}
		lastResult = currentSum.Number
		fmt.Printf("  receive current sum: %d\n", currentSum.Number)
	}

	err = addClient.CloseSend()
	if err != nil {
		fmt.Printf("failed to receive result: %#v\n", err)
		return err
	}

	fmt.Printf("get sum: %d\n", lastResult)

	return nil

}
