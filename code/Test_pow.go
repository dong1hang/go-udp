package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

func main() {

	for i := 0; i < 10000; i++ {
		data := sha256.Sum256([]byte(strconv.Itoa(i)))
		fmt.Printf("%10d,%x\n", i, data)
		fmt.Printf("%s\n", string(data[len(data)-2:]))
		if string(data[len(data)-2:]) == "00" {
			fmt.Println("挖矿成功")
			break
		}
	}
}
