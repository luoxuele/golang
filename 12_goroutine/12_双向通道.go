package main

import (
	"fmt"
	"time"
)

func main() {

	//双向通道
	ch1 := make(chan string)
	go sendData(ch1)

	data := <-ch1
	fmt.Println(data)

	ch1 <- "我是main"

	time.Sleep(time.Second)
	//单向通道

}

func sendData(ch chan string) {
	ch <- "我是田昌"

	data := <-ch
	fmt.Println(data)
}
