# Go Blockchain

A minimal blockchain implementation in Go, built as a hands-on learning project to explore the fundamentals of blockchain technology and the Go programming language.

This project follows the tutorial **[Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)** by Coral Health.

## Overview

The goal of this project is to understand how a basic blockchain works by implementing its core building blocks and exposing them through an HTTP API.

### Features

- **Block structure** — represents each block with an index, timestamp, data, hash, and reference to the previous block.
- **SHA-256 hashing** — generates a hash for each block based on its contents.
- **Chain integrity** — links blocks together through their previous hashes.
- **Block validation** — checks block indexes, previous hashes, and calculated hashes.
- **Longest-chain rule** — provides a basic mechanism for choosing between chains of different lengths.
- **HTTP API** — exposes endpoints to inspect the blockchain and append new blocks.
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
- [Gorilla Mux](https://github.com/gorilla/mux) — HTTP routing.
- [godotenv](https://github.com/joho/godotenv) — environment variable loading.
- [go-spew](https://github.com/davecgh/go-spew) — readable debugging output.

## Getting Started

### Prerequisites

- Go installed on your machine.
- Git.

### Installation

Clone the repository:

```bash
git clone <your-repository-url>
cd <your-repository-directory>
```

Download the project dependencies:

```bash
go mod download
```

### Configuration

Create a `.env` file in the project root:

```dotenv
PORT=8080
```

The application uses the `PORT` environment variable to configure the HTTP server's listening port.

### Run the Application

Start the server with:

```bash
go run .
```

The server should start listening on port `8080`, assuming that is the port configured in your `.env` file.

## API

The application exposes HTTP endpoints to interact with the blockchain.

### Get the Blockchain

```http
GET /
```

Returns the current blockchain as JSON.

### Add a Block

```http
POST /
Content-Type: application/json
```

Example request body:

```json
{
  "BPM": 75
}
```

The server uses the submitted data to generate a new block, calculate its hash, link it to the previous block, and validate it before appending it to the chain.

> Note: The exact endpoints, request fields, and behavior depend on the implementation in this repository.

## How It Works

1. **Genesis block** — initializes the blockchain with its first block.
2. **Block creation** — generates a new block containing a timestamp, an index, and application data.
3. **Hash calculation** — computes a SHA-256 digest from the block's relevant fields.
4. **Chain linking** — stores the previous block's hash in the new block.
5. **Validation** — checks the block's index, previous hash, and calculated hash.
6. **HTTP interaction** — allows clients to retrieve the chain and submit new blocks.

## Limitations

This is a **learning project**, not a production-ready blockchain.

It does not aim to provide a complete consensus protocol, a peer-to-peer network, persistent storage, or production-grade security. The longest-chain rule alone does not provide consensus or protect against malicious participants.

The implementation is intended to make blockchain concepts easier to understand through practical experimentation, rather than to provide a secure distributed ledger.

## References

- **Original tutorial:** [Code your own blockchain in less than 200 lines of Go!](https://mycoralhealth.medium.com/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc) — Coral Health.
- **Go documentation:** [go.dev/doc](https://go.dev/doc/).
- **Gorilla Mux:** [github.com/gorilla/mux](https://github.com/gorilla/mux).

---

Built as a personal learning project while exploring Go and blockchain fundamentals.
