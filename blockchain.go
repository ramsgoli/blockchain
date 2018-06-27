package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Blockchain represents the entire chain
type Blockchain struct {
	Chain               Blocks
	CurrentTransactions []Transaction
	ID                  string
}

// Transaction represents a single transaction in a block
type Transaction struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Amount   int    `json:"amount"`
}

type transactionResponse struct {
	Index int `json:"index"`
}

type mineResponse struct {
	Index        int           `json:"index"`
	Proof        int           `json:"proof"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash [32]byte      `json:"previous_hash"`
}

// NewBlockchain is a constructor-like method which returns an empty blockchain
// with a genesis block
func NewBlockchain() *Blockchain {
	bc := &Blockchain{}

	// adds a genesis block to the chain
	var x [32]byte
	bc.NewBlock(100, x)

	return bc
}

// NewBlock adds a new block to the blockchain
func (bc *Blockchain) NewBlock(proof int, previousHash [32]byte) Block {
	index := len(bc.Chain) + 1
	newBlock := Block{index, time.Now(), bc.CurrentTransactions, proof, previousHash}

	bc.Chain = append(bc.Chain, newBlock)

	// all current transactions now belong to the new block
	bc.CurrentTransactions = nil
	return newBlock
}

// Mine mines a new block
func (bc *Blockchain) Mine(w http.ResponseWriter, r *http.Request) {
	// run proof of work algorithm to get the next proof
	lastBlock := bc.LastBlock()
	lastProof := lastBlock.Proof

	proof := bc.proofOfWork(lastProof)

	// create a new transaction (reward for mining)
	transaction := Transaction{"0", bc.ID, 1}
	bc.CurrentTransactions = append(bc.CurrentTransactions, transaction)

	previousHash := bc.hash(lastBlock)
	block := bc.NewBlock(proof, previousHash)
	res := mineResponse{block.Index, proof, block.Transactions, previousHash}
	json.NewEncoder(w).Encode(res)
}

// GetChain returns an array of blocks on this chain
func (bc *Blockchain) GetChain(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(bc.Chain)
}

// LastBlock returns the last block in the chain
func (bc *Blockchain) LastBlock() Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) hash(block Block) [32]byte {
	bytes, err := json.Marshal(block)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(bytes)
}

// NewTransaction adds a new transaction in the chain
func (bc *Blockchain) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var t Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		panic(err)
	}

	bc.CurrentTransactions = append(bc.CurrentTransactions, t)
	res := transactionResponse{Index: bc.Chain[len(bc.Chain)-1].Index + 1}
	json.NewEncoder(w).Encode(res)
}

/*
	simple proof of work algorithm
	we want to find a hash p such that when hashed with the previous hash,
	we get four leading zeros (Bitcoin uses 18 leading zeros)
	we pretend that the previous hash is the block header
*/
func (bc *Blockchain) proofOfWork(lastProof int) int {
	p := 0
	for !isValidProof(p, lastProof) {
		p++
	}

	return p
}

func isValidProof(nonce int, lastProof int) bool {
	proof := fmt.Sprintf("%d%d", nonce, lastProof)
	byteHash := sha256.Sum256([]byte(proof))
	stringHash := fmt.Sprintf("%s", byteHash)
	return stringHash[:4] == "0000"
}
