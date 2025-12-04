package service

import (
	"context"

	"connectrpc.com/connect"
	pb "github.com/icchy-san/bookstore-app/api/gen/go/github.com/icchy-san/bookstore-app/api/gen/go/bookstore/v1"
)

type BookstoreService interface {
	ListBooks(ctx context.Context, req *connect.Request[pb.ListBooksRequest]) (*connect.Response[pb.ListBooksResponse], error)
}

type bookstoreService struct {
}

func NewBookstoreService() BookstoreService {
	return &bookstoreService{}
}
