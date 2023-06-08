package main

import (
	"context"
	"go-micro-demo/demo-stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	glog "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

type Server struct {
	proto.UnimplementedBroadcastServer
}
func (s *Server) NoneStream(context.Context, *proto.User) (*proto.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoneStream not implemented")
}
func (s *Server) ReturnStream(*proto.User, proto.Broadcast_ReturnStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ReturnStream not implemented")
}
func (s *Server) ReceiveStream(stream proto.Broadcast_ReceiveStreamServer) error {
	return nil

}
func (s *Server) AllStream(proto.Broadcast_AllStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method AllStream not implemented")
}

func main() {
	server := &Server{
	}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	grpcLog.Info("Starting server at port :8080")

	proto.RegisterBroadcastServer(grpcServer, server)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
