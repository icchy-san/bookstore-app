package handler

import (
	"context"

	"github.com/bufbuild/connect-go"
	pb "github.com/icchy-san/bookstore-app/api/gen/go/bookstore/v1"
	"github.com/icchy-san/bookstore-app/api/service"
)

type bookstoreServer struct {
	service service.BookstoreService
}

func (s *bookstoreServer) ListBooks(ctx context.Context, req *connect.Request[pb.ListBooksRequest]) (*connect.Response[pb.ListBooksResponse], error) {
	return s.service.ListBooks(ctx, req)
}

func NewBookstoreServer(service service.BookstoreService) *bookstoreServer {
	return &bookstoreServer{
		service: service,
	}
}
