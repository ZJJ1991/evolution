package block

import (
	"github.com/transformer/evolution/tx"

	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
	"fmt"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Head BlockHead
	Body BlockBody
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
	Transactions []tx.Transaction
}


func (block *Block) Hash() string{
	record := strconv.Itoa(block.Head.Index) + block.Head.Timestamp + strconv.Itoa(block.Head.BPM) + block.Head.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}


// GenerateBlock creates a new block using previous block's hash
func (oldBlock *Block) GenerateBlock(BPM int) *Block {

	var newBlock Block
	t := time.Now()

	newBlock.Head.Index = oldBlock.Head.Index + 1
	newBlock.Head.Timestamp = t.String()
	newBlock.Head.BPM = BPM
	newBlock.Head.PrevHash = oldBlock.Head.Hash
	newBlock.Head.Hash = newBlock.Hash()
	newTxs := make([]tx.Transaction, 1, 5) // initilize the newTxs
	newTxs[0] = tx.Transaction{Nonce:2}
	fmt.Printf("\n newTxs %v", newTxs[0])
	newBlock.Body.Transactions = make([]tx.Transaction, 1, 5)
	copy(newBlock.Body.Transactions[0:], (newTxs)[:len(newTxs)])
	return &newBlock
}

// IsBlockValid makes sure block is valid by checking index, and comparing the hash of the previous block
func (oldBlock *Block) IsBlockValid(newBlock *Block) bool {
	if oldBlock.Head.Index+1 != newBlock.Head.Index {
		return false
	}

	if oldBlock.Head.Hash != newBlock.Head.PrevHash {
		return false
	}

	if 	newBlock.Hash()!= newBlock.Head.Hash {
		return false
	}

	return true
}