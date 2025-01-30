package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func getAllAccounts() ([]Account, error) {
	data := make([]Account, 0)
	err := DB.Select(&data, "select * from accounts")
	if err != nil {
		return nil, fmt.Errorf("Error while getting all accounts: %w", err)
	}
	return data, nil
}

func createAccount(account *Account) (string, error) {
	const query = `
    insert into accounts (
      id, label, balance, treshold
    )
    values (
      :id, :label, :balance, :treshold
    )
  `
	_, err := DB.NamedExec(query, account)
	if err != nil {
		return "", fmt.Errorf("Error while creating new account: %w", err)
	}
	return account.ID, nil
}

func getAccountByID(id string) (*Account, error) {
	const query = `
    select * from accounts where id=$1
  `
	account := &Account{}
	err := DB.Get(account, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("Error while getting account by ID: %w", err)
	}
	return account, nil
}
