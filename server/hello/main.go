package main

import (
	"context"
	"github.com/GoodOneGuy/jiochat/common/rpc"
	"github.com/GoodOneGuy/jiochat/proto/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *service.HelloRequest) (*service.HelloReply, error) {
	trace, _ := metadata.FromIncomingContext(ctx)
	log.Println("get from client:", in.Name, "trace:", trace["trace_id"])
	return &service.HelloReply{Message: "Hello" + in.Name}, nil
}

func main() {

	rpc.StartRpcServer("tcp", ":8000", func(s *grpc.Server) {
		service.RegisterGreeterServer(s, &server{})
	})

}
