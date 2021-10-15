package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	go sendData(ch1)

	//读取通道的数据
	for v :=range ch1{
		
		fmt.Println("读取数据：",v)
		
	}

	fmt.Println("main... over ...")
}

func sendData(ch1 chan int) {
	//发送方：10条数据
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch1 <- i //将i写入通道
	}

	close(ch1)
}
