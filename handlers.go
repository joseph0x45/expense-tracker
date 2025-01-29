package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/oklog/ulid/v2"
)

func serveCSS(w http.ResponseWriter, r *http.Request) {
	f, err := outputCSS.Open("output.css")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer f.Close()
	http.ServeFileFS(w, r, outputCSS, "output.css")
}

func renderHomePage(w http.ResponseWriter, _ *http.Request) {
	template, err := template.ParseFS(viewsFS, "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := getAllAccounts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	template.Execute(w, map[string]interface{}{
		"Accounts": data,
	})
}

func handleAccountCreation(w http.ResponseWriter, r *http.Request) {
	label := r.FormValue("label")
	threshold := r.FormValue("threshold")
	thresholdInt, _ := strconv.Atoi(threshold)
	account := &Account{
		ID:        ulid.Make().String(),
		Label:     label,
		Balance:   0,
		Threshold: thresholdInt,
	}
	_, err := createAccount(account)
	if err != nil {
		if strings.HasSuffix(err.Error(), "UNIQUE constraint failed: accounts.label") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("<div id='error-message' class='text-red-500'>Account label already exists</div>"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<div id='error-message' class='text-red-500'>Something went wrong! Check logs for more info</div>"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
   <a href="/accounts/` + account.ID + `/transactions" class="shadow-sm flex justify-between w-full px-2 py-4 rounded-md border-1 border-gray-300">
			<h1 class="font-semibold">` + account.Label + `</h1>
			<h1 class="` + func() string {
		if account.Balance < account.Threshold {
			return "text-red-500"
		}
		return "text-green-400"
	}() + `">
				` + fmt.Sprintf("%d XOF", account.Balance) + `
			</h1>
		</a> 
  `))
	return
}
