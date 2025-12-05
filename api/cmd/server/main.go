package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/icchy-san/bookstore-app/api/handler"
	"github.com/icchy-san/bookstore-app/api/service"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	bookstoreService := service.NewBookstoreService()

	mux, err := handler.NewBookstoreServiceHandler(bookstoreService)
	if err != nil {
		panic(err)
	}

	log.Printf("Listening on %s:%s\n", host, port)

	if listenErr := http.ListenAndServe(
		fmt.Sprintf("%s:%s", host, port),
		h2c.NewHandler(mux, &http2.Server{}),
	); listenErr != nil {
		if listenErr != http.ErrServerClosed {
			fmt.Fprint(os.Stderr, listenErr)
		}
	}
}
