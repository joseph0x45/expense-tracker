package main

import (
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
