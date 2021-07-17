package main

import (
	"fmt"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"

	"github.com/vsantosalmeida/go-grpc-server/config"
	"github.com/vsantosalmeida/go-grpc-server/pkg/rpc"
	"github.com/vsantosalmeida/go-grpc-server/pkg/stream"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.ServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: time.Second * 10,
		}),
	}

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
