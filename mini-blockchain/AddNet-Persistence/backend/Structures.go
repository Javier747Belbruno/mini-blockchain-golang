package main

type Block struct {
	Index         int          `json:"index"`
	Timestamp     string       `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	Hash          string       `json:"hash"`
	PrevHash      string       `json:"prevhash"`
}

type Transaction struct {
	Hash		string		`json:"hash"`
	Timestamp	string		`json:"timestamp"`
	From   		string 		`json:"from"`
	To     		string 		`json:"to"`
	Amount 		int    		`json:"amount"`
}

var Blockchain []Block

var Transactions []Transaction