package main

import "fmt"

type A struct {
	int
	string
	address string
}
// 字段+类型
//类型 -> 匿名字段-> 等于把类型名用作字段名-> 一个类型只能有一个匿名字段，多了变量名就重复了

func main() {
	// a1 := A{3, "大黄", "地球"}
	a1 := A{}
	a1.int = 2
	a1.string = "大黄"

	fmt.Printf("%#v\n", a1)
}
