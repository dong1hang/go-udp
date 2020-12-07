package main

import "fmt"

//入口

func main() {
	fmt.Println("hello game start ")
	bc := NewBlockchain() //创建区块链
	bc.AddBlock("区块1 pay 10")
	bc.AddBlock("区块2 pay 20")
	bc.AddBlock("区块3 pay 30")

	for _, block := range bc.blocks {
		fmt.Printf("上一块哈希%x\n", block.PrevBlockHash)
		fmt.Printf("数据：%s\n", block.Data)
		fmt.Printf("当前哈希%x", block.Hash)
		fmt.Println()
	}
}
