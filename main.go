package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type config struct {
	TemplatesDir string
	Port         string
	DB           *sqlx.DB
}

//go:embed web/*
var webFS embed.FS

var staticServer http.Handler

func init() {
	publicFS, err := fs.Sub(webFS, "public")
	if err != nil {
		panic(err)
	}
	staticServer = http.FileServer(http.FS(publicFS))
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	envFilePath := fmt.Sprintf("%s.cashflow-logger.env", homeDir)
	godotenv.Load(envFilePath)
	// if err != nil {
	// 	log.Println("Failed to load configuration file at", envFilePath)
	// 	fmt.Println("Run cashflow-logger init to create the file with default values")
	// 	os.Exit(1)
	// }
}

func main() {
	mux := http.NewServeMux()

	mux.Handle(
		"/public/",
		staticServer,
	)

	mux.HandleFunc("/", renderHomePage)

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
