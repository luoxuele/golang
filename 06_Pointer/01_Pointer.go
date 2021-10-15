package main

import "fmt"

func main() {
	//指针存的是另一个变量的内存地址	地址从0开始，所以数值上大于0，相当于uint
	//指针变量的大小	32位系统4字节，64位系统8字节
	//指针变量的类型 	等于指向的变量的类型

	// i := 10
	// pi := &i
	// *pi = 11
	// fmt.Println(i)
	var i int = 10
	var p1 *int = &i
	*p1 = 11
	fmt.Println(i)

	//*T 是指针变量的类型，T可以为int,float64,struct....
	//&	取址，取一个变量的地址
	//* 取值，放在Pointer变量前，整体表示指向的那个变量
	// var p2 *int
	// fmt.Println(p2) //指针默认指向nil
	// fmt.Println(p2 == nil)

	var pp1 **int = &p1
	**pp1 = 12
	fmt.Println(i)

	a := 10
	b := 20
	//a,b = b,a

	// c := a
	// a = b
	// b = c
	swap2(&a,&b)


	fmt.Printf("a = %d , b = %d\n",a,b)
	
}

func swap2(a ,b *int){
	c := *a
	*a = *b
	*b = c
}