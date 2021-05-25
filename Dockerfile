FROM golang:1.15.1-alpine3.12

# install protobuf
RUN apk update \
  && apk add --no-cache git protoc

# install protoc and gRPC
RUN go get -u github.com/golang/protobuf/protoc-gen-go \
  && go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
  && go get -u google.golang.org/grpc

RUN mkdir /app
WORKDIR /app
