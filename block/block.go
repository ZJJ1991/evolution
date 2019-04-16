package block

import (
	"github.com/JChain/evolution/tx"
)
// Block represents each 'item' in the blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// type Block struct {
// 	head BlockHead
// 	body BlockBody
// }

type BlockHead struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

type BlockBody struct {
	transactions []tx.Transaction
}


// Blockchain is a series of validated Blocks
type Blockchain []Block
