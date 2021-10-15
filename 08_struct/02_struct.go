package main

import (
	"fmt"
	"unsafe"
)

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	speciality string
	name string
}

func main() {
	// s := new(Student)
	// s.speciality = "特殊"
	// s.name = "田昌"
	// fmt.Println(s)s
	// i := "ssssss"
	h1 := Human{"田昌",27,65}
	// h1 := Student{}
	fmt.Printf("%T %#v\n",h1,h1)
	fmt.Println(unsafe.Sizeof(h1),"bytes")
}
