package main

import "fmt"

func main() {

	/*
		算数运算符： + - * / % ++ --
	*/

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++
	b++
	fmt.Println(a, b)
	a--
	b--
	fmt.Println(a, b)

}
