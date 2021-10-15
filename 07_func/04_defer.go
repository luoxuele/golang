package main

import "fmt"

func main() {
	//defer 延迟函数，defer会到函数的末尾才执行
	// 一个函数有多个defer,遇到一个就加到栈中，最后执行顺序就变成后进先出，
	myfunc()
}

func myfunc() {
	fmt.Println("我是第1")
	fmt.Println("我是第2")
	defer fmt.Println("我是第3")
	fmt.Println("我是第4")
	defer fmt.Println("我是第5")
	fmt.Println("我是第6")
	defer fmt.Println("我是第7")

	//1,3,6,7,5,3
}
