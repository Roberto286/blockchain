// Package chain implements a consensus-agnostic, concurrency-safe blockchain.
package chain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Block is the minimal view of a block the chain needs; each consensus defines its own concrete type.
type Block interface {
	GetIndex() int
	GetHash() string
	GetPrevHash() string
}

// Generator builds the next block on top of prev; implemented by each consensus.
type Generator func(prev Block, bpm int) Block

// Validator reports whether next is a valid successor of prev; implemented by each consensus.
type Validator func(next, prev Block) bool

var CandidateBlocks = make(chan Block)

func CalculateHash(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
