package main

import "fmt"

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan<- int) //只能写
	// ch3 := make(<-chan int) //只能读

	// ch2 <- 100
	// data := <-ch2

	// data := <-ch3
	// fmt.Println(data)
}

//只能写,单向通道往往是做函数参数的。
func fun1(ch chan<- int) {
	ch <- 100
	fmt.Println("fun1 end ....")
}
