package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 10; i > 0; i-- {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
