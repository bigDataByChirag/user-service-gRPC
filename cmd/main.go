package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
	pb "user-service-gRPC/gen/proto"
	"user-service-gRPC/logger"
)

func main() {
	// SetupSlog initializes the logger.
	logger.SetupSlog()

	// Listen on TCP port 5000.
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		slog.Error("creating user failed", slog.String("Error", err.Error()))
		return
	}

	// NewServer returns the instance of the gRPC server.
	s := grpc.NewServer()

	// Register the user service with the gRPC server.
	pb.RegisterUserServiceServer(s, &userService{})

	// Register reflection service on gRPC server to expose gRPC service for testing with tools like Postman.
	reflection.Register(s)

	// Start serving incoming requests.
	err = s.Serve(listener)
	if err != nil {
		slog.Error("creating user failed", slog.String("Error", err.Error()))
		return
	}
}
