package main
import "fmt"

func main(){
	//整数
	var a int8 = 10
	var b int16
	b = int16(a) //a是int8，需要强转
	fmt.Println(b)

	f1 := 4.13
	var c int
	c = int(f1) //浮点数强转整数取整
	fmt.Println(c)

	b1 := true
	// a = int8(b1)
	// fmt.Println(a) //cannot convert b1 (type bool) to type int8


	/*
	go语言是静态语言，定义，赋值，运算必须类型一致
	语法格式： Type(value)
	常数：在有需要的时候，自动转换 //sum := f1+100 	（100有整数型和浮点型，这里就是浮点型）
	变量： 必须手动强转
	
	*/



	




}