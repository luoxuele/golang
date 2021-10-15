package main

import "fmt"

func main() {
	var a chan int
	fmt.Printf("%T , %v\n", a, a)

	if a == nil{
		fmt.Println("chan int 是空的，不能使用，需要先创建通道")
		a = make(chan int)
		fmt.Println(a)
	}

}
