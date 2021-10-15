package utils

import "fmt"

func Count() {
	fmt.Println("utils.Count().....")
}

// func Count() {
// 	fmt.Println("utils.Count().....")
// }
func init() {
	fmt.Println("utils.init()......utils.go 1")
}

func init() {
	fmt.Println("utils.init()......utils.go 2")
}
