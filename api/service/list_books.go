package service

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/icchy-san/bookstore-app/api/gen/go/github.com/icchy-san/bookstore-app/api/gen/go/bookstore/v1"
)

func (b *bookstoreService) ListBooks(ctx context.Context, req *connect.Request[pb.ListBooksRequest]) (*connect.Response[pb.ListBooksResponse], error) {
	res := connect.NewResponse(&pb.ListBooksResponse{
		Books: []*pb.Book{
			{
				Id:     "1",
				Name:   "Effective Go",
				Author: "Alan A. A. Donovan & Brian W. Kernighan",
			},
		},
	})
	return res, nil
}
