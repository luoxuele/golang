package main

import (
	"fmt"
	"time"
)
//go run -race  03_临界资源.go
func main() {
	a := 1

	go func() {
		a = 2
		fmt.Println("goroutie...", a)
	}()

	a = 3
	time.Sleep(time.Second)
	fmt.Println("main ...", a)
}
