package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	TimeStamp int64

	Data []byte

	PrevBlockHash []byte

	Hash []byte

	Nonce int64
}

func NewBlock(Data string, prevBlockHash []byte) *Block {
	block := &Block{
		TimeStamp:     time.Now().Unix(),
		Data:          []byte(Data),
		PrevBlockHash: prevBlockHash,
		Hash:          nil,
	}

	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	block.Hash = hash

	block.Nonce = nonce

	return block

}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (b *Block) SetHash() {

	timeStamp := []byte(strconv.FormatInt(b.TimeStamp, 10))

	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timeStamp}, []byte{})

	currentHash := sha256.Sum256(headers)

	b.Hash = currentHash[:]

}

func (b *Block) Print() {

	if b == nil {
		return
	}

	fmt.Printf("Drev.hash: %x \n", b.PrevBlockHash)

	fmt.Printf("Data.hash: %s \n", b.Data)

	fmt.Printf("Hash: %x\n", b.Hash)

	fmt.Printf("Nonce: %d\n\n", b.Nonce)

	pow := NewProofOfWork(b)

	fmt.Printf("IsValid: %v\n\n", pow.Validate())

}
