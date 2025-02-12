package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultPort = 8080
	version     = "v1"
)

var (
	Verbose = false
	DB      *sqlx.DB

	//go:embed output.css
	outputCSS embed.FS

	//go:embed *.html
	viewsFS embed.FS
)

func serveContentHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("example.txt")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
}

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	defaultDBFilePath := fmt.Sprintf("%s/.expense_tracker.db", userHomeDir)
	port := flag.Int("port", defaultPort, "Port on which expense tracker should bind to")
	dbFilePath := flag.String("db", defaultDBFilePath, "Path of the SQLite database file")
	flag.BoolVar(&Verbose, "v", false, "Enable verbose mode")
	flag.Parse()
	verbosePrint("Using database at", *dbFilePath)
	if err := checkDBHealth(*dbFilePath); err != nil {
		panic(err)
	}
	log.Println("Starting Expense Tracker", version, " on port", *port)
	DB, err = sqlx.Connect("sqlite3", *dbFilePath)
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()
	DB.SetMaxOpenConns(1)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /output.css", serveCSS)
	mux.HandleFunc("GET /", renderHomePage)
	mux.HandleFunc("POST /api/accounts", handleAccountCreation)
	mux.HandleFunc("GET /accounts/{accountID}/transactions", renderTransactionsPage)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      mux,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
