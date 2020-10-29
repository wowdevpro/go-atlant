package main

import (
	"github.com/wowdevpro/go-atlant/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./cert/server.crt", "./cert/server.key")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	proto.RegisterCsvServiceServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}