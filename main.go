package main

import (
<<<<<<< HEAD
	"flag"
	"fmt"
=======
>>>>>>> parent of d79ef98 (little bit of refactoring)
	"log"
	"net/http"
	"text/template"
	"time"
<<<<<<< HEAD

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sqlx.DB
)

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dbFilePath := fmt.Sprintf("%s/.cashflow.db", userHomeDir)
	_, err = os.Stat(dbFilePath)
	if os.IsNotExist(err) {
		initDB(dbFilePath)
	} else {
		DB, err = sqlx.Connect("sqlite3", dbFilePath)
		if err != nil {
			panic(err)
		}
	}
	log.Println("Connected to database")
=======
)

type film struct {
	Title    string
	Director string
}

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./home.html"))
	tmpl.Execute(w, nil)
>>>>>>> parent of d79ef98 (little bit of refactoring)
}

func main() {
	port := flag.String("port", "8080", "Port on which to launch the app")
	flag.Parse()
	mux := http.NewServeMux()

<<<<<<< HEAD
=======
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	mux.HandleFunc("/", renderHomePage)

>>>>>>> parent of d79ef98 (little bit of refactoring)
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
