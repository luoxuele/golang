package main

import "fmt"

func main() {
	myPrint()
	myPrint2(5)
	myPrint3(1,2,3.3)
	a,b :=myPrint4()
	fmt.Println(a,b)

	//函数是执行特定任务的代码块
	//函数的组成，函数名，形参列表，返回值（列表）
	//函数的定义和调用，func关键字，函数名+()调用
}

func myPrint4()(a int,b float64){
	a = 11
	b = 22.11
	// return 3,4.4
	return a,b
}

func myPrint3(a int,b int,c float64){
	fmt.Println(a,b,c)
}

func myPrint() {
	for i := 10; i > 0; i-- {
		fmt.Printf("%d ",i)
	}
	fmt.Println()
}

func myPrint2(len int) {
	for i := len; i > 0; i-- {
		fmt.Printf("%d ",i)
	}
	fmt.Println()
}
