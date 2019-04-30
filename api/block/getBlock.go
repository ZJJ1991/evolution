package block

import
(
	"encoding/json"
	"net/http"
	"io"
	"github.com/transformer/evolution/chain"
)

type Blocks struct{
	chain *chain.Blockchain
}

// New is used to creates an API instance.
// The real time blockchain is passed to api through this function.
func New(chain *chain.Blockchain) *Blocks {
	return &Blocks{
		chain,
	}
}

// HandleGeBlockchain gets the blockchain data
func (bc *Blocks) HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(bc.chain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}