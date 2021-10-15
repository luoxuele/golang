package main

import "fmt"

func main() {

label:
	fmt.Println("hello")
	goto label
}
