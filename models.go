package main

import "time"

type Transaction struct {
	ID        string    `json:"id" db:"id"`
	Amount    int       `json:"amount" db:"amount"`
	Type      string    `json:"type" db:"flow_type"`
	Purpose   string    `json:"purpose" db:"purpose"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Expected  bool      `json:"expected" db:"expected"`
}

type Account struct {
	ID        string `json:"id" db:"id"`
	Label     string `json:"label" db:"label"`
	Balance   int    `json:"balance" db:"balance"`
	Threshold int    `json:"treshold" db:"treshold"`
}
