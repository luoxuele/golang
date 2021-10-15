package main

import (
	"fmt"
)

func main() {
	s := "hello,世界"
	fmt.Println(len(s))
	// fmt.Println(s[0],s[6])
	// fmt.Println(s[:5])
	// fmt.Println(s[5:])

	// fmt.Println("goodbye "+s[6:])

	//1. string底层数据结构

	//src/runtime/string.go
	// type stringStruct struct {
	// 	str unsafe.Pointer
	// 	len int
	// }

	//src/builtin/builtin.go
	// // string is the set of all strings of 8-bit bytes, conventionally but not
	// // necessarily representing UTF-8-encoded text. A string may be empty, but
	// // not nil. Values of string type are immutable.
	// type string string

	var str string
	fmt.Println(str == "")
	str = "hello 世界"
	fmt.Println(str == "")

	// sstr :=[]uint8{67,68,69,70,71}
	// str = sstr.GetStringBySlice(sstr)
	// fmt.Println(str)

}
