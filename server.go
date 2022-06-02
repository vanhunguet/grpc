package main

import (
	"context"
	pb "github.com/vanhunguet/grpc/gen/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Failed to listen on port 8000: %v", err)
	}

	gprcServer := grpc.NewServer()
	pb.RegisterTestApiServer(gprcServer, &testApiServer{})

	if err := gprcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc serve over port 8080: %v", err)
	}
}
