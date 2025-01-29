package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func verbosePrint(msg ...any) {
	if Verbose {
		log.Println(msg...)
	}
}

func initDB(dbFilePath string) error {
	verbosePrint("Creating new database file at", dbFilePath)
	f, err := os.Create(dbFilePath)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Error while creating database file at: %w", err)
	}
	verbosePrint("Database file created")
	verbosePrint("Attempting connection to database")
	db, err := sqlx.Connect("sqlite3", dbFilePath)
	if err != nil {
		return fmt.Errorf("Error while connecting to database: %w", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)
	verbosePrint("Connected to database")
	const query = `
    create table accounts (
      id text not null primary key,
      label text not null unique,
      balance integer not null default 0
    );
    create table transactions (
      id text not null primary key,
      amount integer not null,
      transaction_type text not null,
      purpose text not null,
      created_at text not null,
      expected integer not null
    );
  `
	verbosePrint("Creating tables")
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("Error while creating tables: %w", err)
	}
  verbosePrint("Tables created! Database ready!")
	return nil
}

// checks if the db file path exists and contains
// all the required tables for expense tracker
func checkDBHealth(dbFilePath string) error {
	verbosePrint("Checking database health")
	verbosePrint("Looking for database file at", dbFilePath)
	_, err := os.Stat(dbFilePath)
	if err == nil {
		verbosePrint("Database file found!")
		verbosePrint("Attempting connection to database")
		db, err := sqlx.Connect("sqlite3", dbFilePath)
		if err != nil {
			return fmt.Errorf("Error while connecting to database: %w", err)
		}
		db.SetMaxOpenConns(1)
		defer db.Close()
		verbosePrint("Connected to database")
		tables := []string{"accounts", "transactions"}
		foundTables := 0
		for _, table := range tables {
			count := 0
			verbosePrint("Looking for table:", table)
			err = db.QueryRow(`
      select count(*) from sqlite_master where type='table'
      and name=$1
      `, table).Scan(&count)
			if err != nil {
				return fmt.Errorf("Error while looking for table '%s': %w", table, err)
			}
			if count != 1 {
				verbosePrint("Table", table, " not found. Database not healthy!")
				break
			}
			verbosePrint("Found table", table)
			foundTables++
		}
		if foundTables == len(tables) {
			verbosePrint("Database healthy!")
			return nil
		}
		verbosePrint("Delete old database file")
		return initDB(dbFilePath)
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("Error while checking database health: %w", err)
	}
	return initDB(dbFilePath)
}
