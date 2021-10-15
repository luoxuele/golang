package main

import "fmt"

type A struct {
	name    string
	age     int
	address string
}
type A2 struct{
	name string
}

func (this *A) say() {
	fmt.Printf("A: i'm %s\n", this.name)
}

type B struct {
	A
	A2
	id string
	// name string
}

func (this *B) say() {
	fmt.Printf("B: I'm %s \n", this.name)
}

func main() {
	//结构体嵌套--》嵌套的结构体就是一普通字段，该怎么用就怎么用
	// 如果嵌套的结构体是匿名的
	// 匿名结构体的字段和方法将得到提升
	// 提升的结果： struct.filed 的时候，会优先使用自己的filed和方法（成员变量，字段，属性）
	// 多个匿名字段有相同的字段，直接报错。

	// a1 := A{}
	// a1.name = "藏三"
	// a1.age = 22
	// fmt.Printf("%#v\n",a1)

	b1 := B{}
	// b1.A.name = "xxx"
	// b1.A.age = 22
	// b1.A.address = "地球"
	b1.name = "呵呵哒"
	b1.age = 22
	b1.address = "火星"
	b1.id = "007"
	b1.say()

	fmt.Printf("%#v\n", b1)

	// a1 := A{"藏三",11,"火星"}
	// a1.say()

	// b1 := B{A{"里斯",22,"地球"},"007"}
	// b1.say()
	// b1.a1.say()
}


