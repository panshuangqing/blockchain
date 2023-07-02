package main

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
	"strconv"
)

const targetBits = 16

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork {

	target := big.NewInt(1)

	target.Lsh(target, uint(256-targetBits))

	return &ProofOfWork{
		block:  block,
		target: target,
	}
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {

	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		pow.IntToHex(pow.block.TimeStamp),
		pow.IntToHex(targetBits),
		pow.IntToHex(nonce),
	}, []byte{})

	return data
}

func (pow *ProofOfWork) IntToHex(target int64) []byte {
	return []byte(strconv.FormatInt(target, 10))
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	hashData := pow.prepareData(pow.block.Nonce)

	hash := sha256.Sum256(hashData)

	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1

}

func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := int64(0)

	for nonce < math.MaxInt {

		hashData := pow.prepareData(nonce)

		hash = sha256.Sum256(hashData)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {

			//fmt.Printf(" find target hash nonce:%v hash:%v \n", nonce, hash)
			return nonce, hash[:]
		}

		nonce++

	}

	return nonce, hash[:]
}
