package main

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/ma-miyazaki/go-grpc-neo4j-example/pb/employee"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
)

func requestAddEmployee(ctx context.Context, conn *grpc.ClientConn, email, lastName, firstName string) error {
	client := employee.NewEmployeeServiceClient(conn)
	reply, err := client.AddEmployee(ctx, &employee.AddEmployeeRequest{
		Email:     email,
		LastName:  lastName,
		FirstName: firstName,
	})
	if err != nil {
		return errors.Wrap(err, "受取り失敗")
	}
	log.Info().Msgf("Reply create employee. [%v]", reply)
	return nil
}

func addEmployee(email, lastName, firstName string) error {
	return doWithConnection(func(conn *grpc.ClientConn) error {
		return doInTimeout(func(ctx context.Context) error {
			return requestAddEmployee(ctx, conn, email, lastName, firstName)
		})
	})
}

func doInTimeout(fx func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		(10 * time.Second),
	)
	defer cancel()
	return fx(ctx)
}

func doWithConnection(fx func(*grpc.ClientConn) error) error {
	address := "localhost:50051"
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return errors.Wrap(err, "コネクションエラー")
	}
	defer conn.Close()
	return fx(conn)
}

func main() {
	if err := addEmployee("test@example.com", "高嶺", "朋樹"); err != nil {
		log.Fatal().Stack().Err(err).Msg("gRPC request error")
	}
}
