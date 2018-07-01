package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
	"github.com/ramsgoli/blockchain/models"
)

func main() {
	blockchain := models.Blockchain{}
	blockchain.NewBlock(100, "")

	uuid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	blockchain.ID = uuid.String()

	router := mux.NewRouter()
	router.HandleFunc("/transactions/new", blockchain.NewTransaction).Methods("POST")
	router.HandleFunc("/chain", blockchain.GetChain).Methods("GET")
	router.HandleFunc("/mine", blockchain.Mine).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
