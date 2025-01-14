package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func initDB(dbPath string) {
	_, err := os.Create(dbPath)
	if err != nil {
		panic(err)
	}
	query := `
    create table if not exists accounts (
      label text not null unique primary key,
      balance integer not null default 0
    );
  `
	DB, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

func logError(err error) {
	log.Println(fmt.Sprintf("[ERROR]: %s", err.Error()))
}
