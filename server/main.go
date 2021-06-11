package main

import (
	"net"

	"github.com/ma-miyazaki/go-grpc-neo4j-example/pb/employee"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/domain/service"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/infrastracture/persistence"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/interface/handler"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/usecase"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

const port = ":50051"

func createEmployeeServer() employee.EmployeeServiceServer {
	repository := persistence.NewEmployeeRepository()
	service := service.NewEmployeeService(repository)
	useCase := usecase.NewEmployeeUseCase(repository, service)
	return handler.NewEmployeeHandler(useCase)
}

func main() {
	defer persistence.CloseNeo4jDriver()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to listen")
		return
	}
	s := grpc.NewServer()
	employee.RegisterEmployeeServiceServer(s, createEmployeeServer())

	if err := s.Serve(lis); err != nil {
		log.Fatal().Stack().Err(err).Msg("failed to serve")
		return
	}
}
