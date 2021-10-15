package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go fun1()
	go fun2()
	
	fmt.Println("main 进入阻塞状态，等待子goroutine结束")
	wg.Wait()
	fmt.Println("main ...解除阻塞")

}

func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("fun1 ...", i)
	}

	wg.Done()

}

func fun2() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("fun2 ...", i)
	}

}
