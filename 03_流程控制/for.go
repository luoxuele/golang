package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			
			if j < 9-i {
				fmt.Printf("  ")
				continue
			}
			fmt.Printf("%d ", j)
		}
		fmt.Println()

	}
}
