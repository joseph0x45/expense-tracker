package main

import (
	"database/sql"
	"errors"
)

func insertAccount(acc *Account) error {
	query := `
    insert into accounts (
      label, balance
    )
    values (
      :label, 0
    )
  `
	_, err := DB.NamedExec(query, acc)
	return err
}

func getAccountByLabel(label string) (*Account, error) {
	acc := &Account{}
	query := "select * from accounts where label=$1"
	err := DB.Get(acc, query, label)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return acc, nil
}
