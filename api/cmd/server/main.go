package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/icchy-san/bookstore-app/api/handler"
	"github.com/icchy-san/bookstore-app/api/service"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	bookstoreService := service.NewBookstoreService()

	mux, err := handler.NewBookstoreServiceHandler(bookstoreService)
	if err != nil {
		panic(err)
	}

	if listenErr := http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); listenErr != nil {
		if listenErr != http.ErrServerClosed {
			fmt.Fprint(os.Stderr, listenErr)
		}
	}
}
