package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
	"github.com/ramsgoli/blockchain/models"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Port argument required")
		os.Exit(2)
	}

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
	router.HandleFunc("/nodes/resolve", blockchain.ResolveConflicts).Methods("GET")
	router.HandleFunc("/nodes/register", blockchain.NewNode).Methods("POST")

	portString := fmt.Sprintf(":%s", os.Args[1])
	log.Fatal(http.ListenAndServe(portString, router))
}
