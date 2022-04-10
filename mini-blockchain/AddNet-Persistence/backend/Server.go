package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
	"sync"
	"time"


	//"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var mutex = &sync.Mutex{}
var port = flag.String("port", "3010", "port to run the server on")


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	flag.Parse()
	
	
	connectDB(*port)
	CreateBlockTable()
	CreateTxTable()


	go func() {
		t := time.Now()
		firstTransaction := Transaction{Hash: "", Timestamp: t.String() ,From: "Alice", To: "Bob", Amount: 500000000000}
		//add Hash to the transaction
		firstTransaction.Hash = calculateHashTx(firstTransaction)
		Transactions = append(Transactions, firstTransaction)
		genesisBlock := Block{}
		genesisBlock = Block{0, "2022-04-04 22:13:39.04866231 -0300 -03 m=+0.167091007",
								 Transactions, 
								"1111111111111111111111111111111111111111111111111111111111111111",
								"0000000000000000000000000000000000000000000000000000000000000000"}

		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		
		
		// save the block to the database
		SaveBlock(genesisBlock)
		// save the tx to the database
		SaveTx(firstTransaction,genesisBlock.Index)

		mutex.Unlock()
	}()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/", HomeLink)

	router.HandleFunc("/api/getblockchain", handleGetBlockchain).Methods("GET")
	router.HandleFunc("/api/sendtransaction", handleWriteBlock).Methods("POST")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	
		
	fmt.Println("now serving on ", *port)
		
	log.Fatal(http.ListenAndServe(":"+*port, handlers.CORS(headers, methods, origins)(router)))

}

