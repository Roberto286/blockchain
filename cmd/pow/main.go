package main

import (
	"blockchain/internal/chain"
	"blockchain/internal/httpserver"
	"blockchain/internal/pow"
	shared "blockchain/internal/shared"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	genesis := pow.Genesis()
	spew.Dump(genesis)
	shared.Mutex.Lock()
	chain.Blockchain = append(chain.Blockchain, genesis)
	shared.Mutex.Unlock()

	log.Fatal(httpserver.Run(pow.GenerateBlock, pow.IsBlockValid))
}
