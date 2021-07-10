package rpc

import (
	"context"
	"log"

	"github.com/vsantosalmeida/go-grpc-server/config"
	"github.com/vsantosalmeida/go-grpc-server/pkg/stream"
	"github.com/vsantosalmeida/go-grpc-server/protobuf"

	"google.golang.org/protobuf/proto"
)

type server struct {
	protobuf.UnimplementedPersonReceiverServer
	producer stream.Producer
}

var errorReply = &protobuf.Reply{Created: false}

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
		return errorReply, err
	}

	recordValue, err := s.producer.ToProtoBytes(messageBytes, config.PersonSubjName)
	if err != nil {
		log.Printf("Failed to send event to stream err: %q", err)
		return errorReply, err
	}

	log.Print("sending person event to stream")

	err = s.producer.Write(recordValue, config.PearsonCreatedTopic)
	if err != nil {
		log.Printf("Failed to send event to stream err: %q", err)
		return errorReply, err
	}

	return &protobuf.Reply{Created: true}, err
}
