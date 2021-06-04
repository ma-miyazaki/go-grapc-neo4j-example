package main

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"

	pb "github.com/ma-miyazaki/go-grpc-neo4j-example/pb/calc"
	"github.com/ma-miyazaki/go-grpc-neo4j-example/pb/employee"

	"google.golang.org/grpc"
)

func requestSum(ctx context.Context, conn *grpc.ClientConn, a, b int32) error {
	client := pb.NewCalcClient(conn)
	sumRequest := pb.SumRequest{
		A: a,
		B: b,
	}
	reply, err := client.Sum(ctx, &sumRequest)
	if err != nil {
		return errors.Wrap(err, "受取り失敗")
	}
	log.Printf("サーバからの受け取り\n %s", reply.GetMessage())
	return nil
}

func sum(a, b int32) error {
	return doWithConnection(func(conn *grpc.ClientConn) error {
		return doInTimeout(func(ctx context.Context) error {
			return requestSum(ctx, conn, a, b)
		})
	})
}

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
	log.Printf("サーバからの受け取り\n %s", reply)
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
		time.Second,
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
	// a := int32(300)
	// b := int32(500)
	// if err := sum(a, b); err != nil {
	// 	log.Fatalf("%v", err)
	// }
	if err := addEmployee("test@example.com", "高嶺", "朋樹"); err != nil {
		log.Fatalf("%v", err)
	}
}
