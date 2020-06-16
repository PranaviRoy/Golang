package connector

import (
	proto "github.com/golang/books_grpc_db/services"

	"google.golang.org/grpc"
)

//Connection is a function to create connection to the server
func Connection() proto.OperationsClient {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return proto.NewOperationsClient(conn) //establishes connection with server
}
