package blockchain

// Blockchain is an array of bocks
type Blockchain struct {
	Blocks []*Block
}

// AddBlock simply adds a new block to the existing blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockchain creates a new Blockchain with genesis Block as the first block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
