package chain

import (
	"errors"
	"slices"
	"sync"
)

var ErrInvalidBlock = errors.New("invalid block")

type Chain struct {
	mu     sync.RWMutex
	blocks []Block
}

var Blockchain []Block
var TempBlocks []Block

// New returns a chain containing only the genesis block.
func New(genesis Block) *Chain {
	return &Chain{blocks: []Block{genesis}}
}

// Add generates a block for bpm with gen, validates it and appends it to the chain.
func (c *Chain) Add(gen Generator, valid Validator, bpm int) (Block, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	prev := c.blocks[len(c.blocks)-1]
	b := gen(prev, bpm)
	if !valid(b, prev) {
		return nil, ErrInvalidBlock
	}
	c.blocks = append(c.blocks, b)
	return b, nil
}

// Replace swaps the chain for blocks if it is longer (longest-chain rule).
func (c *Chain) Replace(blocks []Block) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(blocks) > len(c.blocks) {
		c.blocks = slices.Clone(blocks)
	}
}

// Blocks returns a copy of the current blocks.
func (c *Chain) Blocks() []Block {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return slices.Clone(c.blocks)
}
