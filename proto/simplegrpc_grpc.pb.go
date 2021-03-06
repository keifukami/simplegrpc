// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoClient interface {
	OneEcho(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	MultiEcho(ctx context.Context, in *MultiEchoRequest, opts ...grpc.CallOption) (Echo_MultiEchoClient, error)
}

type echoClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoClient(cc grpc.ClientConnInterface) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) OneEcho(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/proto.Echo/OneEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) MultiEcho(ctx context.Context, in *MultiEchoRequest, opts ...grpc.CallOption) (Echo_MultiEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Echo_ServiceDesc.Streams[0], "/proto.Echo/MultiEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoMultiEchoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Echo_MultiEchoClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type echoMultiEchoClient struct {
	grpc.ClientStream
}

func (x *echoMultiEchoClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServer is the server API for Echo service.
// All implementations must embed UnimplementedEchoServer
// for forward compatibility
type EchoServer interface {
	OneEcho(context.Context, *Message) (*Message, error)
	MultiEcho(*MultiEchoRequest, Echo_MultiEchoServer) error
	mustEmbedUnimplementedEchoServer()
}

// UnimplementedEchoServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (UnimplementedEchoServer) OneEcho(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OneEcho not implemented")
}
func (UnimplementedEchoServer) MultiEcho(*MultiEchoRequest, Echo_MultiEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method MultiEcho not implemented")
}
func (UnimplementedEchoServer) mustEmbedUnimplementedEchoServer() {}

// UnsafeEchoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServer will
// result in compilation errors.
type UnsafeEchoServer interface {
	mustEmbedUnimplementedEchoServer()
}

func RegisterEchoServer(s grpc.ServiceRegistrar, srv EchoServer) {
	s.RegisterService(&Echo_ServiceDesc, srv)
}

func _Echo_OneEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).OneEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Echo/OneEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).OneEcho(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_MultiEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MultiEchoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServer).MultiEcho(m, &echoMultiEchoServer{stream})
}

type Echo_MultiEchoServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type echoMultiEchoServer struct {
	grpc.ServerStream
}

func (x *echoMultiEchoServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

// Echo_ServiceDesc is the grpc.ServiceDesc for Echo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Echo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OneEcho",
			Handler:    _Echo_OneEcho_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MultiEcho",
			Handler:       _Echo_MultiEcho_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/simplegrpc.proto",
}

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Add(ctx context.Context, opts ...grpc.CallOption) (Calculator_AddClient, error)
	AddInteractive(ctx context.Context, opts ...grpc.CallOption) (Calculator_AddInteractiveClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Add(ctx context.Context, opts ...grpc.CallOption) (Calculator_AddClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[0], "/proto.Calculator/Add", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorAddClient{stream}
	return x, nil
}

type Calculator_AddClient interface {
	Send(*Value) error
	CloseAndRecv() (*Value, error)
	grpc.ClientStream
}

type calculatorAddClient struct {
	grpc.ClientStream
}

func (x *calculatorAddClient) Send(m *Value) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorAddClient) CloseAndRecv() (*Value, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Value)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorClient) AddInteractive(ctx context.Context, opts ...grpc.CallOption) (Calculator_AddInteractiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[1], "/proto.Calculator/AddInteractive", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorAddInteractiveClient{stream}
	return x, nil
}

type Calculator_AddInteractiveClient interface {
	Send(*Value) error
	Recv() (*Value, error)
	grpc.ClientStream
}

type calculatorAddInteractiveClient struct {
	grpc.ClientStream
}

func (x *calculatorAddInteractiveClient) Send(m *Value) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorAddInteractiveClient) Recv() (*Value, error) {
	m := new(Value)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	Add(Calculator_AddServer) error
	AddInteractive(Calculator_AddInteractiveServer) error
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) Add(Calculator_AddServer) error {
	return status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorServer) AddInteractive(Calculator_AddInteractiveServer) error {
	return status.Errorf(codes.Unimplemented, "method AddInteractive not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Add_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Add(&calculatorAddServer{stream})
}

type Calculator_AddServer interface {
	SendAndClose(*Value) error
	Recv() (*Value, error)
	grpc.ServerStream
}

type calculatorAddServer struct {
	grpc.ServerStream
}

func (x *calculatorAddServer) SendAndClose(m *Value) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorAddServer) Recv() (*Value, error) {
	m := new(Value)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Calculator_AddInteractive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).AddInteractive(&calculatorAddInteractiveServer{stream})
}

type Calculator_AddInteractiveServer interface {
	Send(*Value) error
	Recv() (*Value, error)
	grpc.ServerStream
}

type calculatorAddInteractiveServer struct {
	grpc.ServerStream
}

func (x *calculatorAddInteractiveServer) Send(m *Value) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorAddInteractiveServer) Recv() (*Value, error) {
	m := new(Value)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Add",
			Handler:       _Calculator_Add_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddInteractive",
			Handler:       _Calculator_AddInteractive_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/simplegrpc.proto",
}
