package main

import (
	"context"
	"fmt"
	"github.com/wowdevpro/go-atlant/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("./cert/server.crt", "localhost")

	conn, err := grpc.Dial("localhost:1443", grpc.WithTransportCredentials(creds))
	//conn, err := grpc.Dial("localhost:5300", grpc.WithTransportCredentials(creds))

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()


	client := proto.NewCsvServiceClient(conn)

	//load 2 files with data
	for _, url := range []string{"https://wowdev.pro/test1.csv", "https://wowdev.pro/test2.csv"} {
		_, err := client.Fetch(context.Background(), &proto.FetchRequest{
			Url: url,
		})

		if err != nil {
			grpclog.Fatalf("%v", err)
		}
	}

	//request with paging and sort params
	res, err := client.List(context.Background(), &proto.ListRequest{
		Paging: &proto.Paging{
			PageSize:   2,
			PageNumber: 1,
		},
		OrderBy: 1,
	})

	if err != nil {
		grpclog.Fatalf("%v", err)
	}

	fmt.Println(res)
}