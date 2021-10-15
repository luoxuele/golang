package main

import (
	"fmt"
	"runtime"
)

//M,G,P,Sched
//G goroutine
// Processor
// Machine 系统线程

func main() {
	fmt.Println(runtime.GOROOT()) //go安装目录
	fmt.Println(runtime.GOOS)	//sting变量，操作系统
	fmt.Println(runtime.NumCPU())	//当前电脑的cpu核数，
	fmt.Println(runtime.GOMAXPROCS(4)) //设置最大使用cpu数量，[1,256] //go1.8之后默认多核

	//runtime.Gosched() //让出时间片
	// runtime.Goexit() //推出当前goroutine

	//Gosched()
	go func(){
		for i:=0;i<10;i++{
			fmt.Println("goroutine...",i)
		}
	}()

	//main
	for i:=0;i<10;i++{
		runtime.Gosched() //让出时间片，让其它的goroutine先执行
		fmt.Println("main...",i)
	}
}
