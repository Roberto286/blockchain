// Package chain implements blocks and a concurrency-safe blockchain.
package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

func calculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.BPM) + b.PrevHash
	sum := sha256.Sum256([]byte(record))
	return hex.EncodeToString(sum[:])
}

func newBlock(prev Block, bpm int) Block {
	b := Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().String(),
		BPM:       bpm,
		PrevHash:  prev.Hash,
	}
	b.Hash = calculateHash(b)
	return b
}

func isValid(next, prev Block) bool {
	return prev.Index+1 == next.Index &&
		prev.Hash == next.PrevHash &&
		calculateHash(next) == next.Hash
}
