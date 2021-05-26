package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ma-miyazaki/go-grpc-neo4j-example/pb/calc"

	"google.golang.org/grpc"
)

const port = ":50051"

type ServerUnary struct {
	pb.UnimplementedCalcServer
}

func (s *ServerUnary) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumReply, error) {
	a := in.GetA()
	b := in.GetB()
	fmt.Println(a, b)
	reply := fmt.Sprintf("%d + %d = %d", a, b, a+b)
	return &pb.SumReply{
		Message: reply,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterCalcServer(s, &ServerUnary{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
