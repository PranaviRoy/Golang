package main

import (
	"strconv"
	"context"
	"net"

	"greeter/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterGreetingServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) Reply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	name, times := request.GetName(), request.GetTimes()

	var result string = ""
	reply := "Annyeong " + name + "!" 
	var i int64 = 0
	for i < times {
		result = result + reply + "(" + strconv.FormatInt(int64(i+1), 10) + ") "
		i++
	}

	return &proto.Response{Reply: result}, nil
}
