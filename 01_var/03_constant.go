package main

import "fmt"

func main() {
	//常数，固定的数值，如300，"abc"
	fmt.Println(100)
	fmt.Println("hello")

	//常量  定义完成初始化后就不能再赋值了
	// const a int //missing value in const declaration
	const PI = 3.14 //常量可以不使用

	// PI = 3.1415926 //cannot assign to PI (declared const)

	//定义一组常量
	const c1, c2, c3 = 100, 3.14, "hehe"
	const (
		MALE   = 0
		FEMALE = 1
		UNKNOW = 3
	)

	// 如果常量没有初始值，默认和上一行一样
	const (
		a int = 100
		b
		c string = "changge"
		d
	)

	fmt.Printf("%T--%d\n", a, a)
	fmt.Printf("%T--%d\n", b, b)
	fmt.Printf("%T--%s\n", c, c)
	fmt.Printf("%T--%s\n", d, d)

	//枚举类型 ：使用常量组作为枚举类型，一组相关数值的数据
	const(
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)

}
