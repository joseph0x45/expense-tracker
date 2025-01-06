package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

type film struct {
	Title    string
	Director string
}

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./home.html"))
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

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
