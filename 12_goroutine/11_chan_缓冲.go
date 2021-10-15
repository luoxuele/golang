package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 非缓存通道： make(chan T)
	//一次发送，一次接收都是阻塞的

	ch1 := make(chan int)
	fmt.Println(len(ch1), cap(ch1))
	// ch1 <- 100  //写入就阻塞了，需要别的goroutine读，才能解除阻塞，否则deadlock

	ch2 := make(chan int, 5)
	ch2 <- 100
	ch2 <- 200
	ch2 <- 300
	ch2 <- 400
	ch2 <- 500
	//ch2 <- 200
	fmt.Println(len(ch2), cap(ch2))

	ch3 := make(chan string, 4)
	go sendData(ch3)
	for {
		v, ok := <-ch3
		if ok{
			fmt.Println("读取的数据是：",v)
		}else{
			fmt.Println("读完了。。。",ok)
			break
		}
	}

	fmt.Println("main...over...")
}

func sendData(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine中写出第%d个数据\n", i)
	}

	close(ch)
}
