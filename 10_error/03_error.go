package main

import (
	"fmt"
	"os"
	"io/fs"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		if ins,ok := err.;ok{
			fmt.Println("1.Op:",ins.Op)
			fmt.Println("2. Path:",ins.Path)
			fmt.Println("3. Error",ins.Error)

		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功")
}
