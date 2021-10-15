package main

import "fmt"

func main() {

	//slice可以看成一个动态数组，底层原理是顺序表，头部是包含了3个字段的结构体
	// type slice struct {
	// 	array unsafe.Pointer
	// 	len   int
	// 	cap   int
	// }

	//slice的定义/声明
	//var identifier []type
	// var s1 []int
	// s1 = make([]int, 10) //len=10, cap=10, 所有元素默认0值
	s1 := make([]int, 10, 20) //len=10, cap=20, 所有元素默认0值
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 0, 5) //len=0,cap=5
	s2 = append(s2, 3, 4, 5)
	_ = append(s2, 11, 22) //append里面有一个局部slice，局部slice的数组并不是外部的数组？
	s2 = append(s2, 666)
	fmt.Println(s2, len(s2), cap(s2))

	//	func append(slice []Type, elems ...Type) []Type
	// 向slice里面追加一个或者多个元素，然后返回一个相同类型的slice

	//copy函数  用第二个slice的内存值覆盖第一个slice的底层数组值
	var numbers []int
    numbers = append(numbers, 0)
    numbers = append(numbers, 1)
    numbers = append(numbers, 2)
    
	// fmt.Println(numbers,len(numbers),cap(numbers))

    numbers2 := make([]int,len(numbers),cap(numbers)*2)
    numbers2 = append(numbers2, 666)
    copy(numbers2,numbers)
    numbers2[0]=999

	fmt.Println(numbers,len(numbers),cap(numbers))
	fmt.Println(numbers2,len(numbers2),cap(numbers2)) //len=3 cap=4*2


}
