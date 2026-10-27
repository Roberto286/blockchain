package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"

	"blockchain/internal/chain"
	"blockchain/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Fatal(server.New(chain.New()).Serve(l))
}
