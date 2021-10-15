package main

import "fmt"

var a int = 100
var b = 200
// c: = 300  //syntax error: non-declaration statement outside function body

func main() {
	var num int = 100
	fmt.Printf("num = %d, &num = %p\n", num, &num)
	num = 200
	fmt.Printf("num = %d, &num = %p\n", num, &num)

	// fmt.Println(num2) //./02_var.go:12:14: undefined: num2

	var name string
	// name = 200  //cannot use 200 (type untyped int) as type string in assignment
	name = "changgle"
	fmt.Println(name)

	// var name int  //name redeclared in this block  重新声明

	// num,name := 300,"田昌" 
	// no new variables on left side of :=
	num,name ,sex:= 300,"田昌" ,"男"
	fmt.Println(num,name,sex)

}
