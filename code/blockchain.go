package main

type BlockChain struct {
	blocks []*Block //一个数组，每个元素都是指针，存储Block区块的地址

}

//增加一个区块
func (blockChain *BlockChain) AddBlock(data string) {
	prevBlock := blockChain.blocks[len(blockChain.blocks)-1] //取出最后一块
	newBlock := NewBlock(data, prevBlock.Hash)               //创建一个区块
	blockChain.blocks = append(blockChain.blocks, newBlock)  //区块链插入新区块
}

//创建一个区块链
func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
