package main

import (
	"Microservice/grpc_ks/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type KingSoft struct {
	// 这个地方以往是空的，可能是生成命令不同
	pb.UnimplementedHelloKSServer
}

func (ks *KingSoft) Welcome(_ context.Context, c *pb.KS) (*pb.KS, error) {
	//s := "welcome " + c.Name
	c.Name += " to kingsoft"
	return c, nil
}

func main() {
	gs := grpc.NewServer()

	pb.RegisterHelloKSServer(gs, new(KingSoft))

	listener, err := net.Listen("tcp", pb.Addr)
	if err != nil {
		log.Fatalln("KS listener error:", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalln("KS listener.Close() error:", err)
		}
	}(listener)

	err = gs.Serve(listener)
	if err != nil {
		log.Fatalln("KS Serve() error:", err)
	}
}
