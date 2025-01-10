package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sqlx.DB

	//go:embed frontend/dist
	frontend embed.FS
)

var buildMode = "prod"

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
	port := flag.String("port", "8080", "Port on which to launch the app")
	flag.Parse()
	mux := http.NewServeMux()
	if buildMode == "dev" {
    log.Println("Running in dev mode")
		frontendURL, err := url.Parse("http://localhost:5173")
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(frontendURL)
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			proxy.ServeHTTP(w, r)
		})
	} else {
		dist, _ := fs.Sub(frontend, "frontend/dist")
		mux.Handle("/", http.FileServer(http.FS(dist)))
	}

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

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
