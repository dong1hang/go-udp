package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//定义区块
type Block struct {
	Timestamp     int64  //时间线，1970/1/1 00：00：00
	Data          []byte //交易数据
	PrevBlockHash []byte //上一块数据的哈希
	Hash          []byte //当前块数据的哈希
}

//设定结构体对象哈希
func (block *Block) SetHash() {
	//处理当前的时间，转化为10进制的字符串，在转化为字节集
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	//叠加要哈希的数据
	headers := bytes.Join([][]byte{block.PrevBlockHash, block.Data, timestamp}, []byte{})
	//计算出哈希地址
	hash := sha256.Sum256(headers)
	block.Hash = hash[:] //设置哈希
}

//创建一个区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	//block 是一个指针，取得一个对象初始化之后的地址
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash() //设置当前哈希
	return block

}

//创建创世区块
func NewGenesisBlock() *Block {

	return NewBlock("天大地大老子最大魔王QQ1234567890", []byte{})
}
