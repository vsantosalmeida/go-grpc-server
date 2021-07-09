FROM golang:1.16-alpine3.13 as builder
RUN apk add --no-cache \
   build-base \
   gcc \
   git \
   pkgconf \
   musl-dev
WORKDIR /go/src/vsantosalmeida/go-grpc-server
RUN export GOPRIVATE=github.com/vsantosalmeida/*
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go get -d -v ./...
RUN make build-server

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/vsantosalmeida/go-grpc-server/bin/go-grpc-server.bin .
EXPOSE 9090
CMD ["./go-grpc-server.bin"]