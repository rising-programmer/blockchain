package main

import (
	"math/big"
	"bytes"
	"math"
	"crypto/sha256"
	"fmt"
)

//目标难度
const targetBits  = 24

/**
 * 工作量证明
 */
type ProofOfWork struct {
	block *Block
	target *big.Int
}

/**
 * 创建pow对象
 */
func NewProofOfWork(block *Block) *ProofOfWork{
	//根据难度值算出实际目标值
	target := big.NewInt(1)
	//向左位移256-targetBits位(也就是常说的前面多少个零)
	target.Lsh(target,uint(256-targetBits))
	pow := &ProofOfWork{block,target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte{
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			[]byte(pow.block.Data),
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(nonce),
		},
		[]byte{},
	)
	return data
}

/**
 * 挖矿
 */
func (pow *ProofOfWork) run() (int,[]byte){
	hashInt := big.Int{}
	nonce := 0
	hash := [32]byte{}
	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < math.MaxInt64{
		data := pow.prepareData(int64(nonce))
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		fmt.Printf("\r%x", hash)
		if hashInt.Cmp(pow.target) == -1{
			break
		}else{
			nonce++
		}
	}
	fmt.Print("\n\n")
	return nonce,hash[:]
}

/**
 * 验证
 */
func (pow * ProofOfWork) IsValidate() bool {
	hashInt := big.Int{}
	data := pow.prepareData(int64(pow.block.Nonce))
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}


