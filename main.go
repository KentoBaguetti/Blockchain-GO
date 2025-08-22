package main

import (
	"fmt"
)

func main() {
	fmt.Println("Init main")

	bc := NewBlockChain()

	bc.AddBlock("Kentaro")
	bc.AddBlock("Barnes")

	for _, block := range bc.blocks {

		fmt.Printf("Prev Hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data %x\n", block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
		fmt.Println()

	}

}