package main

import (
	"blockchain/internal/chain"
	"blockchain/internal/httpserver"
	shared "blockchain/internal/shared"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := chain.Block{}
		genesisBlock = chain.Block{
			Index:      0,
			Timestamp:  t.String(),
			BPM:        0,
			Hash:       chain.CalculateHash(genesisBlock),
			PrevHash:   "",
			Difficulty: shared.Difficulty,
			Nonce:      "",
		}
		spew.Dump(genesisBlock)
		shared.Mutex.Lock()
		chain.Blockchain = append(chain.Blockchain, genesisBlock)
		shared.Mutex.Unlock()
	}()
	log.Fatal(httpserver.Run())
}
