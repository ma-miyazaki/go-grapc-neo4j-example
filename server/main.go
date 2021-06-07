package main

import (
	"net"

	"github.com/ma-miyazaki/go-grpc-neo4j-example/pb/employee"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/infrastracture/persistence"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/interface/handler"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/server/usecase"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

const port = ":50051"

func createEmployeeServer() employee.EmployeeServiceServer {
	employeeRepository := persistence.NewEmployeeRepository()
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepository)
	return handler.NewEmployeeHandler(employeeUseCase)
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
