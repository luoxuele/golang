package main

import "fmt"

type Dog struct {
	name string
	age  int
}

//方法就是把形参里面的 对象（的指针）变量    定义在函数名前面 ，然后取个名字叫接收者，接受者可以是命名类型或者结构体类型的一个值或者是一个指针
// 不同类型的方法可以相同

func (this *Dog) say(){
	fmt.Printf("I'm %s, 旺旺\n",this.name)
}
func say(this *Dog){
	fmt.Printf("I'm %s, 旺旺\n",this.name)
}

func main() {

	d1 := Dog{"大黄", 3}
	//fmt.Println(d1)
	d1.say()

	d2 := Dog{"小花",2}
	d2.say()

	// say(&d1)
	// say(&d2)

}
