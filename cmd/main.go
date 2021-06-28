package main

import (
	"fmt"
	"go-grpc-server/config"
	"go-grpc-server/pkg/rpc"
	"go-grpc-server/pkg/stream"
	"go-grpc-server/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	p, err := stream.NewKafkaProducer()
	if err != nil {
		log.Fatalf("failed to create a kafka producer: %v", err)
	}
	defer p.Close()

	svc := rpc.NewServer(p)

	protobuf.RegisterPersonReceiverServer(grpcServer, svc)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
