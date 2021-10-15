package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rwMutex *sync.RWMutex

func main() {
	rwMutex = new(sync.RWMutex)

	wg.Add(4)
	go readData(1)
	go readData(2)
	go writeData(3)
	go writeData(4)

	wg.Wait()
	fmt.Println("main ... over....")
}

func writeData(i int){
	defer wg.Done()
	fmt.Println(i,"准备写： write start ...")

	rwMutex.Lock()
	fmt.Println(i,"正在写 ：writing ...")
	time.Sleep(time.Second)
	rwMutex.Unlock()

	fmt.Println(i,"写完成： write ..over...")
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "准备读： read start...")

	rwMutex.RLock()
	fmt.Println(i, "正在读取数据： reading...")
	time.Sleep(time.Second)
	rwMutex.RUnlock()

	fmt.Println(i, "读结束：read over ....")
}
