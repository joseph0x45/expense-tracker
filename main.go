package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type config struct {
	Port         string
	DB           *sqlx.DB
}

//go:embed web/*
var templateFS embed.FS

//go:embed web/public/*
var publicFS embed.FS

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	envFilePath := fmt.Sprintf("%s.cashflow-logger.env", homeDir)
	godotenv.Load(envFilePath)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/public/output.css/{$}", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFileFS(w, r, publicFS, "output.css")
	})

	mux.HandleFunc("/home", renderHomePage)

	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
		Handler:      mux,
	}

	log.Println("Server listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
