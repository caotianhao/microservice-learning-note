package main

import (
	"Microservice/grpc_ks/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.NewClient(pb.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("gRPC client error:", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln("gRPC conn.Close() error:", err)
		}
	}(conn)

	kstest := pb.KS{}
	kstest.Name = "cth"
	cli := pb.NewHelloKSClient(conn)

	out, err := cli.Welcome(context.TODO(), &kstest)
	if err != nil {
		log.Fatalln("gRPC Welcome error:", err)
	}

	fmt.Println("kingsoft success!", out)
}
