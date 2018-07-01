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
