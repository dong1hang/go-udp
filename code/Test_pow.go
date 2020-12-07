package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()                //当前时间
	for i := 0; i < 10000000000; i++ { //循环挖矿
		data := sha256.Sum256([]byte(strconv.Itoa(i)))
		fmt.Printf("%10d,%x\n", i, data)
		fmt.Printf("%s\n", string(data[len(data)-1:]))
		if string(data[len(data)-3:]) == "000" {
			usedtime := time.Since(start)
			fmt.Println("挖矿成功", usedtime)
			break
		}
	}
}
