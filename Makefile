BINARY_PATH=./bin
BINARY_NAME=$(BINARY_PATH)/go-grpc-server.bin
VERSION=1.0.0

clean:
	@ rm -rf bin/*

build-server:
	@ echo " ---         BUILDING        --- "
	@ $(MAKE) clean
	@ go build -a -v -tags musl -ldflags "-s -w -X main.version=$(VERSION)" -o $(BINARY_NAME) cmd/main.go
	@ echo " ---      FINISH BUILD       --- "

build-server-docker:
	@ docker build --no-cache -t larolman/go-grpc-server .

push-server-docker-image:
	@ docker login
	@ docker push $(DOCKER_REPO)/go-grpc-server:latest

start-kafka:
	@ docker-compose up -d

stop-kafka:
	@ docker-compose down

create-topics:
	@ bash +x ./create-topics.sh

compile-protobuf:
	@ protoc --proto_path=protobuf --go_out=protobuf --go-grpc_out=protobuf person.proto