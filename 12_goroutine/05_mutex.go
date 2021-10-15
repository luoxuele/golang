package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//全局变量，表示票
var ticket int = 10
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	wg.Add(4)
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait()
}

func saleTickets(name string) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出： ", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "售馨，没有票了")
			break
		}

		mutex.Unlock()

	}
}
