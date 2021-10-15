package main

import "fmt"

func main() {

	//函数的可以没有返回值，也可以有一个或者多个返回值
	fmt.Println(myfunc())
	fmt.Println(myfunc2())
	fmt.Println(myfunc3())

}

func myfunc() (int, int, string) {
	return 2, 3, "hello"
}

func myfunc2() (a, b int, str string) {
	a = 10
	b = 20
	str = "hello golang"

	return a, b, str
}

func myfunc3() (a int) {
	a = 20
	return
}
