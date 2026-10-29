// Package pow implements proof-of-work block generation.
package pow

import (
	"fmt"
	"strings"
	"time"

	"blockchain/internal/chain"
)

const Difficulty = 1

func isHashValid(hash string, difficulty int) bool {
	return strings.HasPrefix(hash, strings.Repeat("0", difficulty))
}

// GenerateBlock mines a block on top of prevBlock; it satisfies chain.Generator.
func GenerateBlock(prevBlock chain.Block, BPM int) chain.Block {
	b := chain.Block{
		Index:      prevBlock.Index + 1,
		Timestamp:  time.Now().String(),
		BPM:        BPM,
		Difficulty: Difficulty,
		PrevHash:   prevBlock.Hash,
	}
	b.Hash = chain.CalculateBlockHash(b)

	for i := 0; ; i++ {
		b.Nonce = fmt.Sprintf("%x", i)
		if !isHashValid(chain.CalculateBlockHash(b), b.Difficulty) {
			fmt.Println(chain.CalculateBlockHash(b), " do more work!")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(chain.CalculateBlockHash(b), " work done!")
		b.Hash = chain.CalculateBlockHash(b)
		break
	}
	return b
}
