package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps the basic structure of a block
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash will be the way we are going to set a hash for our blocks
// 1. We set the time in which the block was created
// 2. The main par of the hash will consist of of the prvious hash, the data and the current time
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates and returns a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates the genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
