package main

import "fmt"

type Dog struct {
	name string
	age  int
}

func (this *Dog) say() {
	fmt.Printf("I'm %s, 旺旺\n", this.name)
}

type Erha struct {
	Dog
}

func main() {

	d1 := Dog{"大黄", 3}
	d1.say()

	d2 := Dog{"小花", 2}
	d2.say()

	// e1 := Erha{d: Dog{"小二", 1}}
	e1 := Erha{Dog{"小二", 1}}
	// e1 := Erha{{"小二", 1}}
	// fmt.Println(e1)
	e1.say()

	//结构体的匿名字段实现了一个方法，那么这个结构体就继承了匿名结构体的方法

}
