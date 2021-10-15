package main

import "fmt"

func main() {
	//iota: 特殊的常量，可以被编译器自动修改（赋值）的常量

	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c)

	const (
		d = iota
		e
		f = 33
		g
		h = iota  //每定义一个常量，iota加1，等于iota就是下标
	)
	fmt.Println(d, e, f, g,h)
}
