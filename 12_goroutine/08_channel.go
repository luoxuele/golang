package main

import "fmt"

func main() {
	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {

		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine ...", i)
		}
		//循环结束，向通道写数据
		ch1 <- true
		fmt.Println("子goroutine over ...")
	}()

	data := <-ch1 //阻塞等待子routine向ch1写数据，
	fmt.Println("main .... data->", data)
	fmt.Println("main ...over ...")

}
