package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type RegisterFunc func(*grpc.Server)

func StartRpcServer(network, addr string, regFunc RegisterFunc) {
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalln("listen tcp err:", err)
		return
	}

	s := grpc.NewServer()
	regFunc(s)
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalln("start server err", err)
		return
	}
}
