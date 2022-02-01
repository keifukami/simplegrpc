package main

import (
	"fmt"
	"github.com/keifukami/simplegrpc/pkg/service/calculator"
	"github.com/keifukami/simplegrpc/pkg/service/echo"
	pb "github.com/keifukami/simplegrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	var err error

	fmt.Println("prepare server...")

	var listener net.Listener
	listener, err = net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	echoServer := echo.NewEchoServer()
	pb.RegisterEchoServer(grpcServer, echoServer)

	calculatorServer := calculator.NewCalculatorServer()
	pb.RegisterCalculatorServer(grpcServer, calculatorServer)

	reflection.Register(grpcServer)

	fmt.Println("start server")
	err = grpcServer.Serve(listener)
	fmt.Println("server finished")
	if err != nil {
		fmt.Println("server returned error")
		panic(err)
	}
	fmt.Println("server successfully returned")

}
