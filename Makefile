BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/go-grpc-server.bin
VERSION=1.0.0

clean:
	@ rm -rf bin/*

build-api:
	@ echo " ---         BUILDING        --- "
	@ go build -ldflags "-s -w -X main.version=$(VERSION)" -o $(BINARY_NAME) cmd/main.go
	@ echo " ---      FINISH BUILD       --- "

start-kafka:
	@ docker-compose up -d

stop-kafka:
	@ docker-compose down

create-topics:
	@ bash +x ./create-topics.sh

compile-protobuf:
	@ protoc --proto_path=protobuf --go_out=protobuf --go-grpc_out=protobuf person.proto