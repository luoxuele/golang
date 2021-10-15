package main

import "fmt"

func main() {
	// var a, b, c, d int
	// var a, b, c, d int = 11,22,33,44
	// fmt.Printf("%T  %v\n", a, a)
	// fmt.Println(a,b,c,d)

	// var a,b,c,d = 1,2,3,4
	// var a bool
	// var b string
	// var c float64
	// var d int
	// a,b,c,d  := false,"hehe",3.3,100
	// fmt.Printf("%T, %T, %T, %T\n", a, b,c,d)
	// fmt.Println(a,b,c,d)

	var (
		a  = 20
		b  = "呵呵"
		c  = true
	)
	fmt.Println(a,b,c)
}
