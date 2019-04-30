package tx
import (
	"math/big"
	"github.com/transformer/evolution/common"
)

type actionBody struct {
	To    *common.Address `rlp:"nil"`
	Value *big.Int
	Data  []byte
}

// Action is a basic represeatation of the conventional transaction.
// A transformer chain transaction contains multiple actions.
type Action struct {
	body actionBody
}