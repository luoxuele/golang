package main

import "fmt"

func main() {
	var arr1 [10]int
	fmt.Println(arr1)

	//var arr2 [10]int = [10]int{1,2,3,4,5}
	// var arr2  = [10]int{1,2,3,4,5,6}
	arr2 := [10]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr2)

	arr3 := [...]string{"hello", "world", "golang", "田昌"}
	// fmt.Printf("arr3 type = %T\n",arr3)
	// fmt.Println("arr3: ",arr3)

	// fmt.Println(arr3[3])
	for i := 0; i < len(arr3); i++ {
		fmt.Println("arr3[", i, "] = ", arr3[i])
	}

	//for range
	// for i,v := range arr3{
	// 	fmt.Printf("arr3[%d] = %s\n",i,v)
	// }
	for _, v := range arr3 {
		fmt.Println(v)
	}

}
