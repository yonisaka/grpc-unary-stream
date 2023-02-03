package main

import (
	"database/sql"
	"fmt"
	"grpc-unary-stream/files/pb"
	"grpc-unary-stream/internal/repositories"
	"grpc-unary-stream/internal/ucase"
	"net"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	conn := registerDB()
	defer conn.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8081))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	repoNote := repositories.NewNoteRepository(conn)
	ucaseNote := ucase.NewNoteUsecase(repoNote)

	pb.RegisterNoteServiceServer(server, ucaseNote)
	reflection.Register(server)
	fmt.Println("Server is running on port 8081")

	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}

func registerDB() *sql.DB {
	strConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", "root", "password", "localhost", "3306", "grpc_us")

	param := url.Values{}
	param.Add("parseTime", "True")
	// param.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", strConn, param.Encode())

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Connected to database")
	return conn
}
