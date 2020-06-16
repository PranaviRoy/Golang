package main

import (
	"fmt"
	"log"
	"net"

	"github.com/golang/books_grpc_db/database"
	proto "github.com/golang/books_grpc_db/services"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	//"google.golang.org/protobuf/types/known/emptypb"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterOperationsServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) AddBook(ctx context.Context, request *proto.Book) (*proto.Response, error) {
	db := database.Conn()

	insForm, err := db.Prepare("INSERT INTO books(name, author, publicationyear, isbn) VALUES( ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	res, err := insForm.Exec(request.GetName(), request.GetAuthor(), request.GetPublicationyear(), request.GetIsbn())
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	request.Id = int64(lastID)
	log.Printf("ID = %d, affected = %d\n", lastID, rowCnt)
	log.Println("INSERT: Id: ", request.Id, insForm)

	defer db.Close()
	return &proto.Response{Result: "Added book successfully!"}, nil
}

func (s *server) FetchBook(ctx context.Context, request *proto.Book) (*proto.Book, error) {
	db := database.Conn()
	defer db.Close()

	book := proto.Book{}
	err := db.QueryRow("select * FROM books where id=?", request.GetId()).Scan(&book.Id, &book.Name, &book.Author, &book.Publicationyear, &book.Isbn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Request successful!")
		fmt.Println(&book)
	}
	return &book, nil
}
