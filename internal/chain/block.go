// Package chain implements blocks and a concurrency-safe blockchain.
package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Generator builds the next block on top of prev; implemented by each consensus.
type Generator func(prev Block, bpm int) Block

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

var CandidateBlocks = make(chan Block)

func CalculateHash(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func CalculateBlockHash(b Block) string {

	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.BPM) + b.PrevHash + b.Nonce

	return CalculateHash(record)
}

func b(prev Block, bpm int) Block {
	b := Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().String(),
		BPM:       bpm,
		PrevHash:  prev.Hash,
	}
	b.Hash = CalculateBlockHash(b)
	return b
}

func IsBlockValid(next, prev Block) bool {
	return prev.Index+1 == next.Index &&
		prev.Hash == next.PrevHash &&
		CalculateBlockHash(next) == next.Hash
}
