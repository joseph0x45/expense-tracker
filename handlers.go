package main

import (
	"encoding/json"
	"net/http"
)

func createAccount(w http.ResponseWriter, r *http.Request) {
	payload := &struct {
		Label string `json:"label"`
	}{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	dbAccount, err := getAccountByLabel(payload.Label)
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if dbAccount != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	err = insertAccount(&Account{
		Label: payload.Label,
	})
	if err != nil {
		logError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
