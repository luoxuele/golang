package main

import "fmt"

func main() {
	//语法结构
	//for init; condition; post { }
	// for condition { }
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	i := 10
	for i > 0 {
		fmt.Printf("%d ", i)
		i--

	}
	fmt.Println()

	// 两种死循环的写法
	// for {
	// 	fmt.Println("for")
	// }
	// for {
	// 	fmt.Println("for")
	// }

	s1 := []int{1, 2, 3, 43, 46, 42, 99, 11}
	for index, value := range s1 {
		fmt.Printf("s1[%d] = %d\n", index, value)
	}

}
