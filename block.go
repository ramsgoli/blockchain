package main

import "time"

// Block represents a single block in the chain
type Block struct {
	Index        int           `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        int           `json:"proof"` // proof given by POW algorithm
	PreviousHash [32]byte      `json:"previous_hash"`
}

type Blocks []Block
