package rpc

import (
	"log"

	"github.com/vsantosalmeida/go-grpc-server/protobuf"

	"google.golang.org/grpc"
)

func NewRpcClient(opts []grpc.DialOption, serverAddr string) (protobuf.PersonReceiverClient, error) {
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Printf("failed to dial to grpc server: %v", err)
		return nil, err
	}
	defer conn.Close()

	return protobuf.NewPersonReceiverClient(conn), nil
}
