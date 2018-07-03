package models

type transactionResponse struct {
	Index int `json:"index"`
}

type mineResponse struct {
	Index        int           `json:"index"`
	Proof        int           `json:"proof"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previous_hash"`
}

func (bc *Blockchain) isValidChain() bool {
	previousBlock := bc.Chain[0]
	currentIndex := 1

	for currentIndex < len(bc.Chain) {
		block := bc.Chain[currentIndex]
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
