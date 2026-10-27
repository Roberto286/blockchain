# Go Blockchain

A minimal blockchain implementation in Go, built as a hands-on learning project to explore the fundamentals of blockchain technology and the Go programming language.

This project follows the tutorials by Coral Health:

1. **[Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)** — blocks, hashing, validation, HTTP API.
2. **[Part 2: Networking — Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/part-2-networking-code-your-own-blockchain-in-less-than-200-lines-of-go-17fe1dad46e1)** — TCP networking (current implementation).

## Overview

The goal of this project is to understand how a basic blockchain works by implementing its core building blocks and exposing them through an HTTP API.

### Features

- **Block structure** — represents each block with an index, timestamp, data, hash, and reference to the previous block.
- **SHA-256 hashing** — generates a hash for each block based on its contents.
- **Chain integrity** — links blocks together through their previous hashes.
- **Block validation** — checks block indexes, previous hashes, and calculated hashes.
- **Longest-chain rule** — provides a basic mechanism for choosing between chains of different lengths.
- **TCP server** — accepts multiple clients; each one submits a BPM value over a raw TCP connection to append a new block.
- **Periodic broadcast** — every 30 seconds each connection receives the current chain as JSON.
- **Environment configuration** — loads configuration from a `.env` file.

## Project Goals

This repository is an ongoing learning project.

It started as a hands-on implementation of the fundamentals covered in the original tutorial, but the goal is to progressively expand it as I learn more about blockchain technology and its underlying concepts.

Future improvements may include:

- Consensus algorithms and mechanisms.
- Peer-to-peer networking and distributed nodes.
- Cryptographic primitives and digital signatures.
- Transaction processing and validation.
- Block and transaction persistence.
- Other blockchain concepts explored along the way.

Rather than aiming to build a production-ready blockchain from the start, the focus is on **learning by building, experimenting, and incrementally improving the implementation** as my understanding of blockchain evolves.

## Tech Stack

- [Go](https://go.dev/)
- [godotenv](https://github.com/joho/godotenv) — environment variable loading.
- [go-spew](https://github.com/davecgh/go-spew) — readable debugging output.

## Getting Started

### Prerequisites

- Go installed on your machine.
- Git.

### Installation

Clone the repository:

```bash
git clone https://github.com/Roberto286/blockchain.git 
cd blockchain 
```

Download the project dependencies:

```bash
go mod download
```

### Configuration

Create a `.env` file in the project root:

```dotenv
ADDR=8080
```

The application uses the `ADDR` environment variable to configure the TCP server's listening port.

### Run the Application

Start the server with:

```bash
go run ./cmd/blockchain
```

The server should start listening on port `8080`, assuming that is the port configured in your `.env` file.

## Usage (TCP)

Connect with netcat (port from `ADDR`):

```bash
nc localhost 8080
```

Type a number (the BPM) and press Enter to append a block. The server dumps the updated chain to its console, and every 30 seconds each client receives the full chain as JSON.

> The HTTP API from part 1 (Gorilla Mux, `GET /` and `POST /`) was replaced by the TCP server in part 2.

## Project Structure

```
cmd/blockchain/   entry point: loads config, starts the TCP server
internal/chain/   Block, hashing, validation, concurrency-safe Chain
internal/server/  TCP connection handling and periodic broadcast
```

## How It Works

1. **Genesis block** — initializes the blockchain with its first block.
2. **Block creation** — generates a new block containing a timestamp, an index, and application data.
3. **Hash calculation** — computes a SHA-256 digest from the block's relevant fields.
4. **Chain linking** — stores the previous block's hash in the new block.
5. **Validation** — checks the block's index, previous hash, and calculated hash.
6. **TCP interaction** — clients connect over TCP, submit BPM values, and periodically receive the chain.

## Limitations

This is a **learning project**, not a production-ready blockchain.

It does not aim to provide a complete consensus protocol, a peer-to-peer network, persistent storage, or production-grade security. The longest-chain rule alone does not provide consensus or protect against malicious participants.

The implementation is intended to make blockchain concepts easier to understand through practical experimentation, rather than to provide a secure distributed ledger.

## References

- **Original tutorial:** [Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc) — Coral Health.
- **Go documentation:** [go.dev/doc](https://go.dev/doc/).
- **Part 2 (networking):** [Coral Health on Medium](https://mycoralhealth.medium.com/part-2-networking-code-your-own-blockchain-in-less-than-200-lines-of-go-17fe1dad46e1).

---

Built as a personal learning project while exploring Go and blockchain fundamentals.
