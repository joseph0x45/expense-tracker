package main

type Account struct {
	Label   string `json:"label" db:"label"`
	Balance int    `json:"balance" db:"balance"`
}
