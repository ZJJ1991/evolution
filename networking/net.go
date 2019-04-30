package network

import (
	"io"
	"net"
	"log"
	"sync"
	"bufio"
	"strconv"
	// "github.com/joho/godotenv"
	"github.com/transformer/evolution/chain"
	blk "github.com/transformer/evolution/block"
)

// NetBlocks is the block representation transmitted in the network.
type NetBlocks struct{
 	bcServer chan []blk.Block
 	mutex  *sync.Mutex
}

var Blockchain = chain.NewChain()

// make sure the chain we're checking is longer than the current blockchain
func (nb *NetBlocks) replaceChain(newBlocks []blk.Block) {
	nb.mutex = &sync.Mutex{}
	nb.mutex.Lock()
	if len(newBlocks) > len(*Blockchain) {
		*Blockchain = newBlocks
	}
	nb.mutex.Unlock()
}


func handleConn(conn net.Conn, comingBlock *blk.Block) {

	defer conn.Close()

	io.WriteString(conn, "Enter a new BPM:")

	scanner := bufio.NewScanner(conn)

	// take in BPM from stdin and add it to blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {
			bpm, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			newBlock := comingBlock.GenerateBlock(bpm)
			if err != nil {
				log.Println(err)
				continue
			}
			if comingBlock.IsBlockValid(newBlock) {
				newBlockchain := append(Blockchain, newBlock)
				replaceChain(newBlockchain)
			}

			bcServer <- Blockchain
			io.WriteString(conn, "\nEnter a new BPM:")
		}
	}()

	// simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			mutex.Lock()
			output, err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			mutex.Unlock()
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(Blockchain)
	}

}

