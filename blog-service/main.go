package main

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"os"
	"time"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()

	if err != nil {
		fmt.Println("an error happen when starting the server: ", err)
		os.Exit(1)
	}
}
