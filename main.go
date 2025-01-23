package main

import (
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
