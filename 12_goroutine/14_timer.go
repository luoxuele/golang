package main

import (
	"fmt"
	"time"
)

func main() {
	// timer := time.NewTimer(time.Second * 3)
	// fmt.Printf("%T\n", timer)
	// fmt.Println(time.Now())

	// ch2 := timer.C
	// fmt.Println(<-ch2)

	timer2 := time.NewTimer(time.Second * 5)

	go func() {
		<-timer2.C
		fmt.Println("Timer2 结束了")
	}()

	time.Sleep(time.Second)
	flag := timer2.Stop()
	if(flag){
		fmt.Println("timer2 停止了")
	}

	fmt.Println("main ...over...")
}
