package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	blockchain := NewBlockchain()

	router := mux.NewRouter()
	router.HandleFunc("/transactions/new", blockchain.NewTransaction).Methods("POST")
	router.HandleFunc("/chain", blockchain.GetChain).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
