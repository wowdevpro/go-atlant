package main

import (
	"context"
	"github.com/wowdevpro/go-atlant/proto"
)

type server struct{}

func (s *server) Fetch(ctx context.Context, request *proto.FetchRequest) (response *proto.FetchResponse, err error) {
	data, err := readCSVFromUrl(request.GetUrl())

	if err != nil {
		return nil, err
	}

	err = saveCsvData(data)
	if err != nil {
		return nil, err
	}

	return &proto.FetchResponse{
		Message: "Done",
	}, nil
}

func (s *server) List(ctx context.Context, request *proto.ListRequest) (response *proto.ListResponse, err error) {
	paging := request.GetPaging()

	products, err := getProducts(request.GetOrderBy().String(), paging.GetPageSize(), paging.GetPageNumber())

	if err != nil {
		return nil, err
	}

	return &proto.ListResponse{
		Product: products,
	}, nil
}
