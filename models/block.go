package models

import "time"

type Block struct {
	Index        int           `json:"index"`
	Timestamp    time.Time     `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Proof        int           `json:"proof"` // proof given by POW algorithm
	PreviousHash string        `json:"previous_hash"`
}

type Blocks []Block
