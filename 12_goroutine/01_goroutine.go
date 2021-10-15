package main

import (
	"fmt"
	"time"
)

func main() {

	go printNum()
	for i := 1; i < 10; i++ {
		fmt.Printf("main 中打印字母%c\n", i+65)
		time.Sleep(time.Second)

	}

	fmt.Println("main over")
}

func printNum() {
	for i := 1; i < 10; i++ {
		fmt.Printf("子goroutine中打印数字：%d\n", i)
		time.Sleep(time.Second)
	}
}
