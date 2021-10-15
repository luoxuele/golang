package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//go中的接口 是一组方法签名，某个类型实现了这个接口的所有方法，称为实现接口
	// go语言中，接口和类型的实现关系是 非侵入式
	//java:  class Mouse implements USB{}

	m1 := Mouse{"罗技"}
	// f1 := FlashDisk{"爱国者"}

	// testInterface(&m1)
	// testInterface(f1)
	var u1 USB
	u1 = &m1

	fmt.Println(unsafe.Sizeof(u1))
	fmt.Printf("%#v\n",u1)

}

//定义接口
type USB interface {
	start()
	end()
}

//实现类
type Mouse struct {
	name string
}

func (this *Mouse) start() {
	fmt.Println(this.name, "鼠标准备就绪。开始工作...")
}
func (this *Mouse) end() {
	fmt.Println(this.name, "工作完成，可以安全退出...")
}

type FlashDisk struct {
	name string
}

func (this FlashDisk) start() {
	fmt.Println(this.name, "U盘开始工作")
}
func (this FlashDisk) end() {
	fmt.Println(this.name, "u盘结束工作，可以拔了")
}

func testInterface(this USB) {
	this.start()
	this.end()
}
