package main

import (
	"net/http"
	"strings"
)

func transactions(w http.ResponseWriter, r *http.Request) {

}

func mine(w http.ResponseWriter, r *http.Request) {

}

func chain(w http.ResponseWriter, r *http.Request) {

}

func newConnection(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func main() {
	blockchain := NewBlockchain()

	http.HandleFunc("/transactions/new", transactions)
	http.HandleFunc("/mine", mine)
	http.HandleFunc("/chain", chain)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
