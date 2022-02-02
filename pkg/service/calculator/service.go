package calculator

import (
	"fmt"
	pb "github.com/keifukami/simplegrpc/proto"
	"google.golang.org/grpc/metadata"
	"io"
	"strings"
)

func NewCalculatorServer() pb.CalculatorServer {
	return &calculatorServer{}
}

type calculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (cs *calculatorServer) Add(stream pb.Calculator_AddServer) error {

	var err error

	var headers metadata.MD
	headers, _ = metadata.FromOutgoingContext(stream.Context())
	logSessionInfo("proto.Calculator/Add", headers)

	var sum int32 = 0
	var operand *pb.Value

	for {

		operand, err = stream.Recv()

		if err == io.EOF {
			err = stream.SendAndClose(&pb.Value{
				Number: sum,
			})
			if err != nil {
				fmt.Printf("send or close failed: %#v", err)
			}
			return err
		}

		if err != nil {
			fmt.Printf("receive failed: %#v", err)
			return err
		}

		sum += operand.Number

	}

}

func (cs *calculatorServer) AddInteractive(stream pb.Calculator_AddInteractiveServer) error {

	var err error

	var headers metadata.MD
	headers, _ = metadata.FromOutgoingContext(stream.Context())
	logSessionInfo("proto.Calculator/AddInteractive", headers)

	var sum int32 = 0
	var operand *pb.Value

	for {

		operand, err = stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			fmt.Printf("receive failed: %#v", err)
			return err
		}

		sum += operand.Number
		err = stream.Send(&pb.Value{
			Number: sum,
		})
		if err != nil {
			fmt.Printf("send failed: %#v", err)
		}

	}

}

func logSessionInfo(rpcName string, md metadata.MD) {
	fmt.Printf("[DEBUG] %s called\n", rpcName)

	for name, values := range md {
		fmt.Printf("[DEBUG]   header name: %s, values: %s.\n", name, strings.Join(values, "; "))
	}
}
