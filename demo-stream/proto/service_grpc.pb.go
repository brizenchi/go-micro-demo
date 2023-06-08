// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/service.proto

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

// BroadcastClient is the client API for Broadcast service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcastClient interface {
	// rpc
	NoneStream(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	// 服务端流
	ReturnStream(ctx context.Context, in *User, opts ...grpc.CallOption) (Broadcast_ReturnStreamClient, error)
	// 客户端流
	ReceiveStream(ctx context.Context, opts ...grpc.CallOption) (Broadcast_ReceiveStreamClient, error)
	// 双向流
	AllStream(ctx context.Context, opts ...grpc.CallOption) (Broadcast_AllStreamClient, error)
}

type broadcastClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcastClient(cc grpc.ClientConnInterface) BroadcastClient {
	return &broadcastClient{cc}
}

func (c *broadcastClient) NoneStream(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.Broadcast/noneStream", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastClient) ReturnStream(ctx context.Context, in *User, opts ...grpc.CallOption) (Broadcast_ReturnStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Broadcast_ServiceDesc.Streams[0], "/proto.Broadcast/returnStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &broadcastReturnStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Broadcast_ReturnStreamClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type broadcastReturnStreamClient struct {
	grpc.ClientStream
}

func (x *broadcastReturnStreamClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *broadcastClient) ReceiveStream(ctx context.Context, opts ...grpc.CallOption) (Broadcast_ReceiveStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Broadcast_ServiceDesc.Streams[1], "/proto.Broadcast/receiveStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &broadcastReceiveStreamClient{stream}
	return x, nil
}

type Broadcast_ReceiveStreamClient interface {
	Send(*User) error
	CloseAndRecv() (*User, error)
	grpc.ClientStream
}

type broadcastReceiveStreamClient struct {
	grpc.ClientStream
}

func (x *broadcastReceiveStreamClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *broadcastReceiveStreamClient) CloseAndRecv() (*User, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *broadcastClient) AllStream(ctx context.Context, opts ...grpc.CallOption) (Broadcast_AllStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Broadcast_ServiceDesc.Streams[2], "/proto.Broadcast/allStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &broadcastAllStreamClient{stream}
	return x, nil
}

type Broadcast_AllStreamClient interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ClientStream
}

type broadcastAllStreamClient struct {
	grpc.ClientStream
}

func (x *broadcastAllStreamClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *broadcastAllStreamClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BroadcastServer is the server API for Broadcast service.
// All implementations must embed UnimplementedBroadcastServer
// for forward compatibility
type BroadcastServer interface {
	// rpc
	NoneStream(context.Context, *User) (*User, error)
	// 服务端流
	ReturnStream(*User, Broadcast_ReturnStreamServer) error
	// 客户端流
	ReceiveStream(Broadcast_ReceiveStreamServer) error
	// 双向流
	AllStream(Broadcast_AllStreamServer) error
	mustEmbedUnimplementedBroadcastServer()
}

// UnimplementedBroadcastServer must be embedded to have forward compatible implementations.
type UnimplementedBroadcastServer struct {
}

func (UnimplementedBroadcastServer) NoneStream(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoneStream not implemented")
}
func (UnimplementedBroadcastServer) ReturnStream(*User, Broadcast_ReturnStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReturnStream not implemented")
}
func (UnimplementedBroadcastServer) ReceiveStream(Broadcast_ReceiveStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveStream not implemented")
}
func (UnimplementedBroadcastServer) AllStream(Broadcast_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}
func (UnimplementedBroadcastServer) mustEmbedUnimplementedBroadcastServer() {}

// UnsafeBroadcastServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcastServer will
// result in compilation errors.
type UnsafeBroadcastServer interface {
	mustEmbedUnimplementedBroadcastServer()
}

func RegisterBroadcastServer(s grpc.ServiceRegistrar, srv BroadcastServer) {
	s.RegisterService(&Broadcast_ServiceDesc, srv)
}

func _Broadcast_NoneStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastServer).NoneStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Broadcast/noneStream",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastServer).NoneStream(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcast_ReturnStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BroadcastServer).ReturnStream(m, &broadcastReturnStreamServer{stream})
}

type Broadcast_ReturnStreamServer interface {
	Send(*User) error
	grpc.ServerStream
}

type broadcastReturnStreamServer struct {
	grpc.ServerStream
}

func (x *broadcastReturnStreamServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _Broadcast_ReceiveStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BroadcastServer).ReceiveStream(&broadcastReceiveStreamServer{stream})
}

type Broadcast_ReceiveStreamServer interface {
	SendAndClose(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type broadcastReceiveStreamServer struct {
	grpc.ServerStream
}

func (x *broadcastReceiveStreamServer) SendAndClose(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *broadcastReceiveStreamServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Broadcast_AllStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BroadcastServer).AllStream(&broadcastAllStreamServer{stream})
}

type Broadcast_AllStreamServer interface {
	Send(*User) error
	Recv() (*User, error)
	grpc.ServerStream
}

type broadcastAllStreamServer struct {
	grpc.ServerStream
}

func (x *broadcastAllStreamServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *broadcastAllStreamServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Broadcast_ServiceDesc is the grpc.ServiceDesc for Broadcast service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broadcast_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Broadcast",
	HandlerType: (*BroadcastServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "noneStream",
			Handler:    _Broadcast_NoneStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "returnStream",
			Handler:       _Broadcast_ReturnStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "receiveStream",
			Handler:       _Broadcast_ReceiveStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "allStream",
			Handler:       _Broadcast_AllStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
