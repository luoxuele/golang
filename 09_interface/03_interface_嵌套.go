package main

import "fmt"

func main() {
	var cat A = Cat{}
	cat.testa()
	// cat.testb()
	// cat.testc()

}

type A interface {
	testa()
}
type B interface {
	testb()
}
type C interface {
	A
	B
	testc()
}

type Cat struct {
}

func (c Cat) testa() {
	fmt.Println("test a ...")
}

func (c Cat) testb(){
	fmt.Println("test b ...")
}

func (c Cat) testc(){
	fmt.Println("test c ...")
}
