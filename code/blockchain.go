package main

type BlockChain struct {
	blocks []*Block //一个数组，每个元素都是指针，存储Block区块的地址

}

//增加一个区块
func (blocks *BlockChain) AddBlock(data string) {

}

//创建一个区块链
func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
