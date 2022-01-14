package main

import (
	"context"
	"net"

	pb "../gen/proto"
	"google.golang.org/grpc"
)

type testApiServe struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServe) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *testApiServe) Getuser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {
	lister, err := net.Listen("tpc", "localhost:8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestApiServer(grpcServer, &testApiServe{})
	grpcServer.Serve(lister)
}
