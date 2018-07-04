package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type transactionResponse struct {
	Index int `json:"index"`
}

type mineResponse struct {
	Index        int           `json:"index"`
	Proof        int           `json:"proof"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previous_hash"`
}

type resolve struct {
	Length  int    `json:"length"`
	Message string `json:"message"`
}

func (bc *Blockchain) isValidChain(chain Blocks) bool {
	previousBlock := chain[0]
	currentIndex := 1

	for currentIndex < len(chain) {
		block := chain[currentIndex]
		if bc.hash(previousBlock) != block.PreviousHash {
			return false
		}

		if !isValidProof(previousBlock.Proof, block.Proof) {
			return false
		}

		previousBlock = block
		currentIndex++
	}

	return true
}

func (bc *Blockchain) ResolveConflicts(w http.ResponseWriter, r *http.Request) {
	neighbors := bc.Nodes
	maxLength := len(bc.Chain)
	var errors []string
	var newChain Blocks

	for _, neighbor := range neighbors {
		chainURL := fmt.Sprintf("%s/chain", neighbor)
		res, err := http.Get(chainURL)
		if err != nil {
			fmt.Println("Could not connec to " + chainURL)
			errors = append(errors, "Could not connect to "+neighbor)
			continue
		}
		defer res.Body.Close()

		var chain Blocks
		if jsonErr := json.NewDecoder(res.Body).Decode(&chain); jsonErr != nil {
			fmt.Println("Could not parse json")
			errors = append(errors, "Could not parse json")
			continue
		}

		if bc.isValidChain(chain) && len(chain) > maxLength {
			maxLength = len(chain)
			newChain = chain
		}
	}

	var s resolve
	if len(newChain) > 0 {
		bc.Chain = newChain
		s = resolve{maxLength, "Found new chain"}
	} else {
		s = resolve{maxLength, "This nodes chain is the longest"}
	}
	jsonS, _ := json.Marshal(s)
	w.Write(jsonS)
}
