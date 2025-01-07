package main

import (
	"fmt"
)

func getAccounts() ([]Account, error) {
	const query = "select * from accounts"
	data := make([]Account, 0)
	err := DB.Select(&data, query)
	if err != nil {
		return nil, fmt.Errorf("Error while getting accounts: %w", err)
	}
	return data, nil
}
