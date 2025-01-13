package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sqlx.DB
)

var buildMode = "prod"

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./home.html"))
	tmpl.Execute(w, nil)
}

func main() {
	port := flag.String("port", "8080", "Port on which to launch the app")
	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("/", renderHomePage)

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", *port),
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
