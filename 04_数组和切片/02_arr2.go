package main

import "fmt"

func main() {
	//数组是值类型
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr2[0] = 5

	fmt.Println(arr1, arr2)

	//数组的大小是类型的一部分，[5]int 和 [6]int是不同的类型，不能相互赋值
	// var arr3 [6]int
	// arr3 = arr2 //cannot use arr2 (variable of type [5]int) as [6]int value in assignment


	
}
