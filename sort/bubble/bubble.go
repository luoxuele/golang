package main

import "fmt"

func sort(arr []int) {
	len := len(arr)
	// fmt.Println(len)
	for i := 0; i < len-1; i++ {
		//fmt.Printf("%d : ", i)
		for j := 0; j < len-i-1; j++ {
			fmt.Printf("%d ", j)
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		//fmt.Println()
	}
}

func main() {
	arr := []int{10, 23, 234, 6, 36, 83, 15, 90, 22, 55}
	sort(arr)
	fmt.Println(arr)
}
