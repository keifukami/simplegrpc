package echo

import (
	"context"
	"fmt"
	pb "github.com/keifukami/simplegrpc/proto"
	"time"
)

func NewEchoServer() pb.EchoServer {
	return &echoServer{}
}

type echoServer struct {
	pb.UnimplementedEchoServer
}

func (es *echoServer) OneEcho(ctx context.Context, message *pb.Message) (*pb.Message, error) {
	resp := pb.Message{
		SourceName:      message.DestinationName,
		DestinationName: message.SourceName,
		Body:            message.DestinationName,
	}
	return &resp, nil
}

func (es *echoServer) MultiEcho(echoRequest *pb.MultiEchoRequest, stream pb.Echo_MultiEchoServer) error {

	var err error

	var repeats uint32 = 3
	if echoRequest.Repeats != nil {
		repeats = *echoRequest.Repeats
	}

	var delay time.Duration = 0
	if echoRequest.DelayInSeconds != nil {
		delay = ((time.Duration)(*echoRequest.DelayInSeconds)) * time.Second
	}

	echoMessage := pb.Message{
		SourceName:      echoRequest.Message.DestinationName,
		DestinationName: echoRequest.Message.SourceName,
		Body:            echoRequest.Message.Body,
	}

	var ticker *time.Ticker
	if delay > 0 {
		ticker = time.NewTicker(delay)
	}

	var count uint32
	for count = 1; count <= repeats; count++ {

		if delay > 0 {
			<-ticker.C
		}

		echoMessage.Count = &count
		err = stream.Send(&echoMessage)
		if err != nil {
			fmt.Printf("send failed: %#v", err)
			return err
		}

	}
	return nil

}
