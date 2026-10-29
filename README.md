# Go Blockchain

A minimal blockchain implementation in Go, built as a hands-on learning project to explore the fundamentals of blockchain technology and the Go programming language.

This project follows the tutorials by Coral Health:

1. **[Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)** — blocks, hashing, validation, HTTP API.
2. **[Part 2: Networking — Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/part-2-networking-code-your-own-blockchain-in-less-than-200-lines-of-go-17fe1dad46e1)** — TCP networking (implementation replaced by part 3).
3. **[Code your own blockchain mining algorithm in Go!](https://mycoralhealth.medium.com/code-your-own-blockchain-mining-algorithm-in-go-82c6a71aba1f)** — mining / proof-of-work (current implementation).

## Overview

The goal of this project is to understand how a basic blockchain works by implementing its core building blocks and exposing them through an HTTP API.

### Features

- **Block structure** — represents each block with an index, timestamp, data, hash, reference to the previous block, difficulty and nonce.
- **SHA-256 hashing** — generates a hash for each block based on its contents (including the nonce).
- **Chain integrity** — links blocks together through their previous hashes.
- **Block validation** — checks block indexes, previous hashes, and calculated hashes.
- **Mining (proof-of-work)** — a new block is only accepted once its hash starts with `Difficulty` leading zeros; the miner tries increasing nonces until one hashes to a valid block.
- **Longest-chain rule** — provides a basic mechanism for choosing between chains of different lengths.
- **HTTP API** — `GET /` returns the chain as JSON; `POST /` with a BPM value mines and appends a block.
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

The application uses the `ADDR` environment variable to configure the HTTP server's listening port.

### Run the Application

Start the server with:

```bash
go run ./cmd/pow
```

The server should start listening on port `8080`, assuming that is the port configured in your `.env` file.

## Usage (HTTP)

`GET /` returns the current chain as JSON:

```bash
curl http://localhost:8080/
```

`POST /` mines and appends a new block for the given BPM:

```bash
curl -X POST -H 'Content-Type: application/json' \
     -d '{"BPM": 70}' \
     http://localhost:8080/
```

The server responds with the newly mined block (`201 Created`). Note that mining takes a moment: the miner tries nonces one at a time until the hash has the required number of leading zeros.

> The TCP server from part 2 (periodic chain broadcast) was replaced by the HTTP API again in part 3, following the mining tutorial.

## Project Structure

```
cmd/pow/          entry point: loads config, creates the genesis block, starts the HTTP server
internal/chain/   Block, hashing, validation, Generator type (consensus-agnostic)
internal/pow/     proof-of-work block generation (mining) and Difficulty
internal/httpserver/  HTTP API (Gorilla Mux), takes a chain.Generator
internal/shared/  the safety Mutex
```

## How It Works

1. **Genesis block** — initializes the blockchain with its first block (difficulty set, empty previous hash).
2. **Block creation** — generates a new block containing a timestamp, an index, and application data.
3. **Hash calculation** — computes a SHA-256 digest from the block's relevant fields.
4. **Mining (proof-of-work)** — iterates over nonces until the block hash starts with `Difficulty` leading zeros, then records the winning nonce and hash.
5. **Chain linking** — stores the previous block's hash in the new block.
6. **Validation** — checks the block's index, previous hash, and calculated hash.
7. **HTTP interaction** — clients submit BPM values and fetch the chain over the HTTP API.

## Limitations

This is a **learning project**, not a production-ready blockchain.

It does not aim to provide a complete consensus protocol, a peer-to-peer network, persistent storage, or production-grade security. The longest-chain rule alone does not provide consensus or protect against malicious participants.

The implementation is intended to make blockchain concepts easier to understand through practical experimentation, rather than to provide a secure distributed ledger.

## References

- **Original tutorial:** [Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc) — Coral Health.
- **Go documentation:** [go.dev/doc](https://go.dev/doc/).
- **Part 2 (networking):** [Coral Health on Medium](https://mycoralhealth.medium.com/part-2-networking-code-your-own-blockchain-in-less-than-200-lines-of-go-17fe1dad46e1).
- **Part 3 (mining):** [Coral Health on Medium](https://mycoralhealth.medium.com/code-your-own-blockchain-mining-algorithm-in-go-82c6a71aba1f).

---

Built as a personal learning project while exploring Go and blockchain fundamentals.
