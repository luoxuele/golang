package main

import "fmt"

func main() {

	// if语句 语法格式
	if true {
		fmt.Println("if true")
	}

	if true {
		fmt.Println("if else --- true")
	} else {
		fmt.Println("if else --- false")
	}

	if false {
		fmt.Println("else if 1")
	} else if false {
		fmt.Println("else if 2")

	} else if false {
		fmt.Println("else if 3")

	}else{
		fmt.Println("else if 4")
	}


	// if 变体
	// num变量是if语句的局部变量，只在当前if语句有效
	if num:=10 ; num%2 == 0{
		fmt.Println(num," is even")
	}else{
		fmt.Println(num," is odd")
	}

}
