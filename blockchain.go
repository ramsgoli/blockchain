package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// Blockchain represents the entire chain
type Blockchain struct {
	Chain               []*Block
	CurrentTransactions []*Transaction
}

// Block represents a single block in the chain
type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []*Transaction
	Proof        int // proof given by POW algorithm
	PreviousHash string
}

// Transaction represents a single transaction in a block
type Transaction struct {
	Sender   string
	Receiver string
	Amount   int
}

// NewBlockchain is a constructor-like method which returns an empty blockchain
// with a genesis block
func NewBlockchain() *Blockchain {
	bc := &Blockchain{}

	// adds a genesis block to the chain
	bc.NewBlock(100, "")
	return bc
}

// NewBlock adds a new block to the blockchain
func (bc *Blockchain) NewBlock(proof int, previousHash string) *Block {
	index := len(bc.Chain) + 1
	newBlock := &Block{index, time.Now(), bc.CurrentTransactions, proof, previousHash}

	bc.Chain = append(bc.Chain, newBlock)

	// all current transactions now belong to the new block
	bc.CurrentTransactions = nil
	return newBlock
}

// LastBlock returns the last block in the chain
func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

// NewTransaction creates a new transaction in the chain
func (bc *Blockchain) NewTransaction(sender string, receiver string, amount int) int {
	transaction := &Transaction{sender, receiver, amount}
	bc.CurrentTransactions = append(bc.CurrentTransactions, transaction)

	// return index of block that will hold this transaction
	return bc.Chain[len(bc.Chain)-1].Index + 1
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
