// Package pow implements proof-of-work block generation.
package pow

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"blockchain/internal/chain"
)

const Difficulty = 1

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

func (b Block) GetIndex() int       { return b.Index }
func (b Block) GetHash() string     { return b.Hash }
func (b Block) GetPrevHash() string { return b.PrevHash }

// CalculateBlockHash hashes the block fields, nonce included.
func CalculateBlockHash(b Block) string {
	return chain.CalculateHash(strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.BPM) + b.PrevHash + b.Nonce)
}

// Genesis returns the first block of a PoW chain.
func Genesis() Block {
	g := Block{Timestamp: time.Now().String(), Difficulty: Difficulty}
	g.Hash = CalculateBlockHash(g)
	return g
}

// IsBlockValid satisfies chain.Validator.
func IsBlockValid(next, prev chain.Block) bool {
	n, ok := next.(Block)
	if !ok {
		return false
	}
	return prev.GetIndex()+1 == n.Index &&
		prev.GetHash() == n.PrevHash &&
		CalculateBlockHash(n) == n.Hash
}

func isHashValid(hash string, difficulty int) bool {
	return strings.HasPrefix(hash, strings.Repeat("0", difficulty))
}

// GenerateBlock mines a block on top of prevBlock; it satisfies chain.Generator.
func GenerateBlock(prevBlock chain.Block, BPM int) chain.Block {
	b := Block{
		Index:      prevBlock.GetIndex() + 1,
		Timestamp:  time.Now().String(),
		BPM:        BPM,
		Difficulty: Difficulty,
		PrevHash:   prevBlock.GetHash(),
	}
	b.Hash = CalculateBlockHash(b)

	for i := 0; ; i++ {
		b.Nonce = fmt.Sprintf("%x", i)
		if !isHashValid(CalculateBlockHash(b), b.Difficulty) {
			fmt.Println(CalculateBlockHash(b), " do more work!")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(CalculateBlockHash(b), " work done!")
		b.Hash = CalculateBlockHash(b)
		break
	}
	return b
}
