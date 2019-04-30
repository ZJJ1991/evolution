package main

import (
	// "os"
	"github.com/transformer/evolution/block"
	"github.com/transformer/evolution/tx"
	apiblock "github.com/transformer/evolution/api/block"
	"log"
	"net/http"
	"sync"
	"time"
	"github.com/transformer/evolution/chain"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
)

// Message takes incoming JSON payload for writing heart rate
type Message struct {
	BPM int
}

var mutex = &sync.Mutex{}

// Blockchain is a blockchain
var Blockchain = chain.NewChain()


func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	go func() {
		t := time.Now()
		genesisBlock := &block.Block{}
		genesisBlock.Head = block.BlockHead{0, t.String(), 0, genesisBlock.Hash(), ""}
		txs := make([]tx.Transaction, 1)
		txs[0] = tx.Transaction{1}
		genesisBlock.Body = block.BlockBody{Transactions: txs}
		spew.Dump(genesisBlock)

		mutex.Lock()
		*Blockchain = append(*Blockchain, *genesisBlock)
		mutex.Unlock()
	}()
	log.Fatal(run())
}

// web server
func run() error {
	mux := makeMuxRouter()
    // httpPort := os.Getenv("PORT")
    httpPort := "8080"
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

// create handlers
func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", apiblock.New(Blockchain).HandleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", apiblock.New(Blockchain).HandleWriteBlock).Methods("POST")
	return muxRouter
}

