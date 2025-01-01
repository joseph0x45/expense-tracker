package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const insertFlowQuery = `
  insert into flows(
    flow_type, created_at, amount,
    method, planned, purpose
  )
  values(
    :flow_type, :created_at,
    :amount, :method, :planned, :purpose
  )
`

func getDBFilePath() string {
	user_homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error while getting user home dir", err.Error())
		os.Exit(1)
	}
	return fmt.Sprintf("%s/.cashflow.db", user_homedir)
}

func initDB() {
	initQuery := `
    create table if not exists flows (
      id integer primary key,
      flow_type text not null,
      created_at text not null,
      amount integer not null,
      method text not null,
      planned text not null,
      purpose text not null
    )
  `
	db, err := sqlx.Connect("sqlite3", getDBFilePath())
	defer db.Close()
	if err != nil {
		fmt.Println("Failed to connect to database:", err.Error())
		db.Close()
		os.Exit(1)
	}
	_, err = db.Exec(initQuery)
	if err != nil {
		fmt.Println("Error while initializing database:", err.Error())
		db.Close()
		os.Exit(1)
	}
	fmt.Println("Database ready!")
}

func handleReaderError(field string, err error) {
	str := fmt.Sprintf("Error while reading %s: %s", field, err.Error())
	fmt.Println(str)
	os.Exit(1)
}

func getFormattedDateTime() string {
	return time.Now().Format("Monday 02 January 2006 15h04")
}

func trimLineFeed(str *string) {
	*str = strings.TrimSuffix(*str, "\n")
}

func printFlow(f *flow) {
	fmt.Println("##########################")
	fmt.Printf("Flow Type: %s\n", f.FlowType)
	fmt.Printf("Amount: %d\n", f.Amount)
	fmt.Printf("Created At: %s\n", f.CreatedAt)
	fmt.Printf("Method: %s\n", f.Method)
	fmt.Printf("Planned: %s\n", f.Planned)
	fmt.Println("##########################")
}

func getDBConnection() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", getDBFilePath())
	if err != nil {
		fmt.Println("Failed to connect to database:", err.Error())
		return nil
	}
	return db
}

func saveFlow(f *flow) {
	db := getDBConnection()
	if db == nil {
		os.Exit(1)
	}
	_, err := db.NamedExec(insertFlowQuery, f)
	if err != nil {
		fmt.Println("Failed to save flow:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Flow saved successfully!")
}

func getAllFlows() []flow {
	db := getDBConnection()
	defer db.Close()
	data := make([]flow, 0)
	err := db.Select(&data, "select * from flows")
	if err != nil {
		fmt.Println("Error while getting all flows:", err.Error())
		return nil
	}
	return data
}
