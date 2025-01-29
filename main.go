package main

import (
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
)

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
