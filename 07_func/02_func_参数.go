package main

import "fmt"

func main() {
	//形参，函数定义的时候
	//实参，函数调用的时候传给形参的
	myfunc(1, 2, 3, 6, 5, 4)

}

//可变参，其实就是个slice
func myfunc(arg ...int) {
	fmt.Printf("%T\n", arg)
	fmt.Println(arg)
}
