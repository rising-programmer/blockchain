package main

type BlockChain struct {
	blocks []*Block
}

func (blockChain *BlockChain) AddBlock(data string){
	prevBlock := blockChain.blocks[len(blockChain.blocks) - 1]
	block := NewBlock(data,prevBlock.Hash)
	blockChain.blocks = append(blockChain.blocks,block)
}

func NewBlockChain() *BlockChain{
	blockChain := &BlockChain{[]*Block{NewGenesisBlock()}}
	return blockChain
}