package rpc

import (
	"context"
	"github.com/vsantosalmeida/go-grpc-server/config"
	"github.com/vsantosalmeida/go-grpc-server/pkg/stream"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"
	"google.golang.org/protobuf/proto"
	"log"
)

type server struct {
	protobuf.UnimplementedPersonReceiverServer
	producer stream.Producer
}

func NewServer(producer stream.Producer) protobuf.PersonReceiverServer {
	return &server{
		producer: producer,
	}
}

func (s *server) CreateEvent(ctx context.Context, p *protobuf.Person) (*protobuf.Reply, error) {
	log.Print("event received")
	messageBytes, err := proto.Marshal(p)
	if err != nil {
		log.Printf("Failed to serealize person err: %q", err)
		return &protobuf.Reply{Created: false}, err
	}

	recordValue := s.producer.ToProtoBytes(messageBytes, config.PersonSubjName)

	log.Print("sending person event to stream")

	err = s.producer.Write(recordValue, config.PearsonCreatedTopic)
	if err != nil {
		log.Printf("Failed to send event to stream err: %q", err)
		return &protobuf.Reply{Created: false}, err
	}

	return &protobuf.Reply{Created: true}, err
}
