package handler

import (
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/icchy-san/bookstore-app/api/gen/go/bookstore/v1/v1connect"
	"github.com/icchy-san/bookstore-app/api/service"
)

func NewBookstoreServiceHandler(service service.BookstoreService, opts ...connect.HandlerOption) (http.Handler, error) {
	mux := http.NewServeMux()

	server := NewBookstoreServer(service)
	path, handler := v1connect.NewBookstoreServiceHandler(server, opts...)
	mux.Handle(path, handler)

	return mux, nil
}
