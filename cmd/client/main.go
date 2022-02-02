package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "github.com/keifukami/simplegrpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type config struct {
	ServerAddress string
	EnableTLS     bool
	TLS           *tlsOption
}

func (c *config) LoadFromEnv() {

	serverAddress := os.Getenv("SIMPLEGRPC_SERVER_ADDRESS")
	if serverAddress == "" {
		serverAddress = "localhost:50051"
	}
	c.ServerAddress = serverAddress

	enableTLS := os.Getenv("SIMPLEGRPC_ENABLE_TLS")
	c.EnableTLS = enableTLS == "true"

	if c.EnableTLS {
		t := &tlsOption{}
		if t.LoadFromEnv() {
			c.TLS = t
		}
	}

}

func (c *config) Credential() grpc.DialOption {

	var err error

	if !c.EnableTLS {
		return grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	serverName := strings.Split(c.ServerAddress, ":")[0]

	cfg := &tls.Config{
		ServerName: serverName,
	}

	if c.TLS != nil {

		var ca []byte
		ca, err = ioutil.ReadFile(c.TLS.CACertPath)
		if err != nil {
			panic(fmt.Errorf("failed to read CA certificate from %s: %#v", c.TLS.CACertPath, err))
		}

		certPool := x509.NewCertPool()
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			panic(fmt.Errorf("failed to append CA certificate to pool: %#v", err))
		}

		var certs []tls.Certificate
		if c.TLS.Client != nil {
			var clientCerts tls.Certificate
			clientCerts, err = tls.LoadX509KeyPair(c.TLS.Client.CertPath, c.TLS.Client.KeyPath)
			if err != nil {
				panic(fmt.Errorf("failed to load Client cert or key: cert: %s, key: %s: %#v", c.TLS.Client.CertPath, c.TLS.Client.KeyPath, err))
			}
			certs = append(certs, clientCerts)
		}

		cfg.RootCAs = certPool
		cfg.Certificates = certs

	}

	return grpc.WithTransportCredentials(credentials.NewTLS(cfg))

}

type tlsOption struct {
	CACertPath string
	Client     *clientCert
}

func (t *tlsOption) LoadFromEnv() bool {

	caCertPath := os.Getenv("SIMPLEGRPC_TLS_CA_CERT_PATH")
	if caCertPath != "" {

		t.CACertPath = caCertPath

		c := &clientCert{}
		if c.LoadFromEnv() {
			t.Client = c
		}

		return true

	}

	return false

}

type clientCert struct {
	CertPath string
	KeyPath  string
}

func (c *clientCert) LoadFromEnv() bool {

	certPath := os.Getenv("SIMPLEGRPC_TLS_CLIENT_CERT_PATH")
	keyPath := os.Getenv("SIMPLEGRPC_TLS_CLIENT_KEY_PATH")

	if certPath != "" && keyPath != "" {
		c.CertPath = certPath
		c.KeyPath = keyPath
		return true
	}

	return false

}

func main() {

	var err error

	cfg := &config{}
	cfg.LoadFromEnv()

	var opts = []grpc.DialOption{
		cfg.Credential(),
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(cfg.ServerAddress, opts...)
	if err != nil {
		fmt.Printf("failed to dial: %s\n", cfg.ServerAddress)
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
		fmt.Printf("OneEcho failed: %#v, %s\n", err, err)
	}

	fmt.Println("-----")
	fmt.Println("call proto.Echo/MultiEcho")
	err = callMultiEcho(echoClient, "myname", "yourname", "Hello, World!", 4, 2)
	if err != nil {
		fmt.Printf("MultiEcho failed: %#v, %s\n", err, err)
	}

	fmt.Println("-----")
	fmt.Println("call proto.Calculator/Add")
	err = callAdd(calcClient, []int32{10, 5, 200})
	if err != nil {
		fmt.Printf("Add failed: %#v, %s\n", err, err)
	}

	fmt.Println("-----")
	fmt.Println("call proto.Calculator/AddInteractive")
	err = callAddInteractive(calcClient, []int32{10, 5, 200})
	if err != nil {
		fmt.Printf("AddInteractive failed: %#v, %s\n", err, err)
	}

}

func logSessionInfo(header, trailer metadata.MD) {
	fmt.Println("[DEBUG]   headers:")
	for name, values := range header {
		fmt.Printf("[DEBUG]     name: %s, values: %s.\n", name, strings.Join(values, "; "))
	}
	fmt.Println("[DEBUG]   trailers:")
	for name, values := range trailer {
		fmt.Printf("[DEBUG]     name: %s, values: %s.\n", name, strings.Join(values, "; "))
	}
}

func callOneEcho(client pb.EchoClient, myName string, yourName string, message string) error {

	var err error

	fmt.Println("[DEBUG] call proto.Echo/OneEcho")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	msg := &pb.Message{
		SourceName:      myName,
		DestinationName: yourName,
		Body:            message,
	}

	var resp *pb.Message
	var header, trailer metadata.MD
	resp, err = client.OneEcho(ctx, msg, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		return err
	}
	logSessionInfo(header, trailer)
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

	fmt.Println("[DEBUG] call proto.Echo/MultiEcho")

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
	var header, trailer metadata.MD
	header, err = stream.Header()
	if err != nil {
		return err
	}

	var msg *pb.Message
	for {

		msg, err = stream.Recv()

		if err == io.EOF {
			trailer = stream.Trailer()
			logSessionInfo(header, trailer)
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

	fmt.Println("[DEBUG] call proto.Calculator/Add")

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

	var header, trailer metadata.MD
	header, err = addClient.Header()
	if err != nil {
		return err
	}
	trailer = addClient.Trailer()
	logSessionInfo(header, trailer)

	fmt.Printf("get sum: %d\n", sumValue.Number)

	return nil

}

func callAddInteractive(client pb.CalculatorClient, operands []int32) error {

	var err error

	fmt.Println("[DEBUG] call proto.Calculator/AddInteractive")

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

	var header, trailer metadata.MD
	header, err = addClient.Header()
	if err != nil {
		return err
	}
	trailer = addClient.Trailer()
	logSessionInfo(header, trailer)

	fmt.Printf("get sum: %d\n", lastResult)

	return nil

}
