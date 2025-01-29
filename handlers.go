package main

import (
	"log"
	"net/http"
)

func handleGetAllAccounts(w http.ResponseWriter, _ *http.Request) {
	accounts, err := getAllAccounts()
	if err != nil {
		log.Println(err.Error())
		writeError(w, http.StatusInternalServerError, "Something went wrong! Check the logs for more info")
		return
	}
	writeData(w, http.StatusCreated, map[string]interface{}{
		"accounts": accounts,
	})
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) {
}
