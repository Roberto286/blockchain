// Package chain implements blocks and a concurrency-safe blockchain.
package chain

import (
	shared "blockchain/internal/shared"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}

func CalculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp + strconv.Itoa(b.BPM) + b.PrevHash + b.Nonce
	sum := sha256.Sum256([]byte(record))
	return hex.EncodeToString(sum[:])
}

func b(prev Block, bpm int) Block {
	b := Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().String(),
		BPM:       bpm,
		PrevHash:  prev.Hash,
	}
	b.Hash = CalculateHash(b)
	return b
}

func IsBlockValid(next, prev Block) bool {
	return prev.Index+1 == next.Index &&
		prev.Hash == next.PrevHash &&
		CalculateHash(next) == next.Hash
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)

	return strings.HasPrefix(hash, prefix)
}

func GenerateBlock(prevBlock Block, BPM int) Block {

	b := Block{
		Index:      prevBlock.Index + 1,
		Timestamp:  time.Now().String(),
		BPM:        BPM,
		Difficulty: shared.Difficulty,
		PrevHash:   prevBlock.Hash,
	}
	b.Hash = CalculateHash(b)

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		b.Nonce = hex
		if !isHashValid(CalculateHash(b), b.Difficulty) {
			fmt.Println(CalculateHash(b), " do more work!")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(CalculateHash(b), " work done!")
		b.Hash = CalculateHash(b)
		break
	}
	return b
}
