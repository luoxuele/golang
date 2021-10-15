package main

import "fmt"

func main() {
	//数组-》存储一组同类型的数据
	//struct 把相同或不同的数据类型组成一个整体

	type Person struct {
		name string //成员，字段
		age  int
		sex  string
	}

	// var p1 Person
	// p1.name = "田昌"
	// p1.age = 24
	// p1.sex = "男"

	// var p1 Person = Person{"zhangsan",23,"男"}
	// var p1 = Person{"李四",22,"中"}
	// p1 := Person{"李四",22,"中"}

	//初始化结构体
	// p1 := Person{"田昌",14,"男"}
	// p1 := Person{sex:"男",name:"田昌"}
	p1 := new(Person)  //返回的是一个堆内存的地址，p1其实是一个指针
	p1.name = "田昌"
	p1.age = 22
	p1.sex = "男"


	fmt.Printf("%T\t%v\n",p1,p1)
}
