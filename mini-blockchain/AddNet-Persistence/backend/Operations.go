package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " baby blockchain API")
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Tx Transaction

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&Tx); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()

	fmt.Println("Print")
	fmt.Println(len(Transactions))

	// Timestamp Tx
	t := time.Now()
	Tx.Timestamp = t.String()
	// Hash Tx
	Tx.Hash = calculateHashTx(Tx)

	//If list of transactions is with length of 3, create a new block and a new list of transactions 
	if len(Transactions) == 3 {
		newBlock := generateBlock(Blockchain[len(Blockchain)-1], Tx)
		spew.Dump(newBlock)
		if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			Blockchain = append(Blockchain, newBlock)
			Transactions = Transactions[:0]
			fmt.Println("-----------------------------------------")
			spew.Dump(Blockchain)
			fmt.Println(Transactions)
			fmt.Println("-----------------------------------------")
		}	
	} else {
		// Add Tx to the list of transactions of the current block
		Transactions = append(Transactions, Tx)
		// Modify hash of the last block
		Blockchain[len(Blockchain)-1].Hash = calculateHashBlock()
		Blockchain[len(Blockchain)-1].Transactions = Transactions

		// update hash of the last block
		//UpdateHashBlock
	}

	mutex.Unlock()

	//respondWithJSON(w, r, http.StatusCreated, "Hola")

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	//if calculateHash(newBlock) != newBlock.Hash {
		//return false
	//}

	return true
}

func calculateHashTx(tx Transaction) string {
	record := tx.From + tx.To + strconv.Itoa(tx.Amount) + tx.Timestamp
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func calculateHashBlock() string {

	var totalAmount int

	// for each transaction, add the amount to the total amount
	for _, tx := range Transactions {
		totalAmount += tx.Amount
	}

	fmt.Println(totalAmount)
	var block = Blockchain[len(Blockchain)-1]
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(totalAmount) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Tx Transaction) Block {

	var newBlock Block
	var newTransactions []Transaction

	t := time.Now()
	
	transactions := newTransactions
	Transactions = append(transactions, Tx)


	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Transactions = Transactions
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHashBlock()

	return newBlock
}
