package block

import
(
	"encoding/json"
	"net/http"
	// "io"
	// "github.com/transformer/evolution/chain"

	"github.com/davecgh/go-spew/spew"
	"github.com/transformer/evolution/api/utils"
	"fmt"
	"sync"
)

// Message refers to the message in a block.
type Message struct {
	BPM int
}

var mutex = &sync.Mutex{}

// HandleWriteBlock takes JSON payload as an input for heart rate (BPM)
func (bc *Blocks) HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg Message

    decoder := json.NewDecoder(r.Body)
    fmt.Print("v%\n", r)
	if err := decoder.Decode(&msg); err != nil {
        println("bad http request", r.Body)
		utils.RespondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
    }
	defer r.Body.Close()
	mutex.Lock()
	prevBlock := (*bc.chain)[len(*(bc.chain))-1]
	newBlock := (&prevBlock).GenerateBlock(msg.BPM)
	// newBlock := GenerateBlock(&prevBlock, msg.BPM)
	if (&prevBlock).IsBlockValid(newBlock) {
		*bc.chain = append(*bc.chain, *newBlock)
		spew.Dump(bc.chain)
	}
	mutex.Unlock()
	utils.RespondWithJSON(w, r, http.StatusCreated, newBlock)

}

