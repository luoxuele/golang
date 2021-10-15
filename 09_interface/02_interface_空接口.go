package main

import "fmt"

func main() {
	var a1 A = Cat{"Blue"}
	var a2 A = Person{"王二狗", 22}
	var a3 A = "haha"
	var a4 A = 1000
	fmt.Println(a1, a2, a3, a4)

	test1("hehe")
	test2("嘿嘿")

	//空接口没有任何的方法，所以等于所有的类型都实现了空接口
	//fmt包下的Print系类函数
	// func Println(a ...interface{}) (n int, err error)
	// func Printf(format string, a ...interface{}) (n int, err error)
	// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
	//

	//map, key字符串 value任意类型
	map1 := make(map[string] interface{})
	map1["name"] = "田昌"
	map1["age"] = 27
	map1["address"] = "地球"
	fmt.Println(map1)

	//切片
	slice1 := make([]interface{},0,10)
	slice1 = append(slice1, a1,a2,a3,a4,"我新")
	fmt.Println(slice1)

}

func test1(a A) {
	fmt.Println(a)
}

func test2(xxx interface{}) {
	fmt.Println(xxx)
}

type A interface{}

type Cat struct {
	Color string
}

type Person struct {
	name string
	age  int
}
