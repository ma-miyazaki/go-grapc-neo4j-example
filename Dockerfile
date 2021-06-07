FROM golang:1.16.4-alpine3.13

# install protobuf
RUN apk update \
  && apk add --no-cache git protoc

# install protoc and gRPC
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0 \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0 \
  && go install github.com/cosmtrek/air@v1.27.3

RUN mkdir /app
WORKDIR /app
