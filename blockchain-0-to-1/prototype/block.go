package main

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Timestamp int64
	Data string
	PrevHash []byte
	Hash []byte
}

func NewBlock(data string,prevHash []byte) *Block{
	block := &Block{time.Now().Unix(),data,prevHash,[]byte{}}
	block.SetHash()
	return block
}

func (block *Block) SetHash(){
	timeStamp := []byte(strconv.FormatInt(block.Timestamp,10))
	header := bytes.Join([][]byte{timeStamp,[]byte(block.Data),block.PrevHash},[]byte{})
	hash := sha256.Sum256(header)
	block.Hash = hash[:]
}

func NewGensisBlock() * Block{
	block := NewBlock("创始区块",[]byte{})
	return block
}
