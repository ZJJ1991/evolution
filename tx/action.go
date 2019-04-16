package tx
import (
	"math/big"
	"github.com/JChain/evolution/common"
)
type actionBody struct {
	To    *ethcommon.Address `rlp:"nil"`
	Value *big.Int
	Data  []byte
}

// Action is a basic represeatation of the conventional transaction.
// A JChain transaction contains multiple actions.
type Action struct {
	body actionBody
}