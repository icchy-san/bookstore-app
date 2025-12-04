package main

import (
	"fmt"
	"net/http"

	"github.com/icchy-san/bookstore-app/api/gen/go/github.com/icchy-san/bookstore-app/api/gen/go/bookstore/v1/v1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type BookstoreServer struct {
	v1connect.UnimplementedBookstoreServiceHandler
}

func main() {
	bookstore := &BookstoreServer{}

	mux := http.NewServeMux()

	path, handler := v1connect.NewBookstoreServiceHandler(bookstore)
	fmt.Println(path)
	fmt.Println(handler)

	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
