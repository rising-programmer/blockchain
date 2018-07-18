package main

import (
	"time"
)

/**
 * 区块结构
 */
type Block struct {
	TimeStamp int64
	Data string
	PrevHash []byte
	Hash []byte
	Nonce int
}

/**
 * 生成新的区块
 */
func NewBlock(data string,prevHash []byte) *Block{
	block := &Block{time.Now().Unix(),data,prevHash,[]byte{},0}
	pow := NewProofOfWork(block)
	nonce,hash := pow.run()
	block.Hash = hash
	block.Nonce = nonce
	return block
}

/**
 * 生成创世区块
 */
func NewGenesisBlock() * Block{
	block := NewBlock("创始区块",[]byte{})
	return block
}
