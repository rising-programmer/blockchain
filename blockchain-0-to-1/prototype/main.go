package main

import "fmt"

func main(){
	blockChain := NewBlockChain()
	blockChain.AddBlock("第二个区块")
	blockChain.AddBlock("第三个区块")
	for _,block := range blockChain.blocks{
		fmt.Printf("PrevHash: %x\n",block.PrevHash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Println()
	}
}
