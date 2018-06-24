package blockchain

import (
    "time"
)

type Blockchain struct {
    Chain []int
    CurrentTransactions []int
}

type Block struct {
    Index int
    Timestamp time.Time
    Transactions []int
    Proof string
    PreviousHash string
}

type Transaction struct {
    Sender string
    Receiver string
    amount int
}

