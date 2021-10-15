package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	go sendData(ch1)

	//读取通道的数据
	for {
		v, ok := <-ch1
		if ok {
			fmt.Println("读取到数据：", v,ok)
		}else{
			fmt.Println("通道已经关闭了",v,ok)
			break
		}
	}
}

func sendData(ch1 chan int) {
	//发送方：10条数据
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ch1 <- i //将i写入通道
	}

	close(ch1)
}
