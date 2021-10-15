package main

import "fmt"

func main() {
	//variable 变量，固定内存大小的一个别名

	//使用：
	/*
		step1: 变量的声明，也叫定义
		step2: 变量的访问，赋值和取值
	*/

	//第一种，声明一个变量，然后赋值
	var num1 int //实际上会默认赋值，int默认0
	// num1 = 30
	fmt.Println(num1)

	//第二种，声明并赋值
	var num2 int = 33
	fmt.Println(num2)

	//go的特性：静态语言，强类型语言，
	//第二种： 类型推断 省略类型
	var name = "田昌"
	fmt.Printf("类型是: %T, 值是%s\n", name, name)

	//第三种 简短定义，也叫简短声明  省略var关键字，用:=代替
	// 类型推导的另一种写法
	sum := 100
	fmt.Println(sum)

	//多个变量同时定义  相同的变量
	var a, b, c int
	// a, b, c = 3, 2, 1
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	//不同的变量只能用类型推导或简短定义
	// var n1, f1, s1 = 100, 3.14, "Golang"
	n1, f1, s1 := 100, 3.14, "Golang"
	fmt.Println(n1, f1, s1)

	//定义一组变量
	var (
		studentName = "田昌"
		age         = 27
		sex         = "男"
	)
	fmt.Println(studentName, age, sex)
}
