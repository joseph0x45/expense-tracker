package main

import (
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
      id integer not null primary key,
      label text not null,
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
