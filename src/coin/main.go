package main

import (
	"core"
	"fmt"
)

func main() {
	bc := core.NewBlockChain()

	bc.AddBlock("one")
	bc.AddBlock("two")

	for _, block := range bc.Blocks {
		fmt.Printf("prevHash %x \n", block.PrevBlockHash)
		fmt.Printf("data     %s \n", block.Data)
		fmt.Printf("Hash     %x \n", block.Hash)
		fmt.Println()
	}
}
