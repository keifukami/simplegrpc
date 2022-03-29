package main

import (
	"fmt"
	"github.com/keifukami/simplegrpc/pkg/service/calculator"
	"github.com/keifukami/simplegrpc/pkg/service/echo"
	pb "github.com/keifukami/simplegrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	TCPNetworkName = "tcp"
	TCPPortMin = 0
	TCPPortMax = 65535
)
var (
	FactoryDefaultListenAddress = &net.TCPAddr {
		IP: net.IPv4(0, 0, 0, 0),
		Port: 50051,
	}
)

func main() {
	var err error

	fmt.Println("prepare server...")

	cfg := &conf{
		ListenAddress: FactoryDefaultListenAddress,
	}

	var cfgFromEnv *conf
	cfgFromEnv, err = loadServerConfigFromEnv()
	if err != nil {
		panic(err)
	}
	cfg.Override(cfgFromEnv)

	var opts []grpc.ServerOption
	grpcServer := buildGRPCServer(opts)

	fmt.Println("start server")
	err = listenAndServe(cfg, grpcServer)
	fmt.Println("server finished")
	if err != nil {
		fmt.Println("server returned error")
		panic(err)
	}
	fmt.Println("server successfully returned")

}

func buildGRPCServer(opts []grpc.ServerOption) *grpc.Server {

	grpcServer := grpc.NewServer(opts...)

	echoServer := echo.NewEchoServer()
	pb.RegisterEchoServer(grpcServer, echoServer)

	calculatorServer := calculator.NewCalculatorServer()
	pb.RegisterCalculatorServer(grpcServer, calculatorServer)

	reflection.Register(grpcServer)

	return grpcServer

}

func listenAndServe(c *conf, s *grpc.Server) error {

	var err error

	var listener net.Listener
	listener, err = net.Listen(TCPNetworkName, c.ListenAddress.String())
	if err != nil {
		return err
	}

	return s.Serve(listener)

}

type conf struct {
	ListenAddress *net.TCPAddr
}

func loadServerConfigFromEnv() (*conf, error) {

	var err error

	listenAddress := os.Getenv("SIMPLEGRPC_LISTEN_ADDRESS")
	if listenAddress == "" {
		return &conf{ListenAddress: nil}, nil
	}

	ip := net.IPv4(0, 0, 0, 0)
	lastColonIdx := strings.LastIndex(listenAddress, ":")
	if lastColonIdx >= 0 { // otherwise, use 0.0.0.0 as ip part
		ipStr := listenAddress[:lastColonIdx]
		if ipStr != "" { // otherwise, use 0.0.0.0 as ip part
			ip = net.ParseIP(ipStr)
		}
	}

	portStr := listenAddress[lastColonIdx+1:]
	var port int
	port, err = strconv.Atoi(portStr)
	if err != nil{
		return nil, err
	}
	if port < TCPPortMin || TCPPortMax < port {
		return nil, fmt.Errorf("TCP port number out of range: %d", port)
	}

	addr := &net.TCPAddr{
		IP: ip,
		Port: port,
	}

	return &conf{ListenAddress: addr}, nil

}

func (c *conf) Override(other *conf) {
	if other.ListenAddress != nil {
		c.ListenAddress = other.ListenAddress
	}
}
