package main

import (
	"fmt"
	"os"
)

func main() {

	// 1. 路径
	// fileName1 := "/root/git/golang/pkg/os/test/readme.txt"
	// fileName2 := "test/readme.txt"

	// fmt.Println(filepath.IsAbs(fileName1))
	// fmt.Println(filepath.IsAbs(fileName2))

	// fmt.Println(filepath.Abs(fileName1))
	// fmt.Println(filepath.Abs(fileName2))

	// //父目录
	// fmt.Println(path.Join(fileName2,".."))

	//2 . 目录 0777
	// err := os.Mkdir("test/xxx",os.ModePerm)
	// if err != nil{
	// 	fmt.Println("err: ",err)
	// 	return
	// }

	// fmt.Println("mkdir ok")

	//os.MkdirAll("test/aaa/bbb/ccc",os.ModePerm)

	// 3. 创建文件	0666
	// file1, err := os.Create("test/xxx.txt")
	// if err != nil{
	// 	fmt.Println("os.Create err: ",err)
	// 	return
	// }
	// fmt.Println(file1)

	// file2, err := os.Open("test/xxx.txt")
	// if err != nil {
	// 	fmt.Println("os.Open() err: ", err)
	// }
	// fmt.Println(file2)

	// file3, err := os.OpenFile("test/xxx.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	// if err != nil {
	// 	fmt.Println("os.Openfile() err: ", err)
	// 	return
	// }
	// fmt.Println(file3)
	// file3.Close()

	// os.Create()  截断创建
	// os.Open() 只读打开
	// os.OpenFile()
	// file.Close()



	// const (
	// 	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	// 	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// 	// The remaining values may be or'ed in to control behavior.
	// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	// )
	

	//5. 删除文件
	//Remove 	RemoveAll(递归删除)
	err := os.Remove("test/xxx")
	if err != nil{
		fmt.Println("os.Remove() err: ",err)
		return
	}



}
