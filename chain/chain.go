package chain

import (
	"github.com/transformer/evolution/block"
)

// Blockchain is a series of validated Blocks
type Blockchain []block.Block

// NewChain initilizes the local real time global blockchain variable,
// and the real time blockchain objects used in other packages
// must be obtained from here.
func NewChain() *Blockchain{
	bc := make(Blockchain, 0)
	return &bc
}

// GetCurrentChain fetches the current blockchain in the local node.
func (bc *Blockchain) GetCurrentChain() *Blockchain{
	return bc
}

// Insertchain inserts the blocks into the current local node.
func (bc *Blockchain) Insertchain(chain *Blockchain) *Blockchain {
	if (len(*bc) < len(*chain)){
		bc = chain
		return bc
	}
	return bc
}

// CompareChain derives the 
// func (bc *Fullchain) CompareChain(chain *Fullchain) *Fullchain{
// }