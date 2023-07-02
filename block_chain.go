package main

type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{
		blocks: []*Block{NewGenesisBlock()},
	}
}

func (bc *BlockChain) AddBlock(Data string) {

	// 创世块
	if len(bc.blocks) == 0 {

		return
	}

	prev := bc.blocks[len(bc.blocks)-1]

	newBlock := NewBlock(Data, prev.Hash)

	bc.blocks = append(bc.blocks, newBlock)

}

func (bc *BlockChain) Print() {

	if bc == nil || len(bc.blocks) == 0 {
		return
	}

	for ii := 0; ii < len(bc.blocks); ii++ {
		bc.blocks[ii].Print()
	}
}
