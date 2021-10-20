package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "xxx.txt"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// b1 := bufio.NewReader(file)

	// p := make([]byte,1024)
	// n1,err := b1.Read(p)
	// fmt.Println(n1)
	// fmt.Println(string(p[:n1]))

	// data,flag,err := b1.ReadLine()
	// fmt.Println(err)
	// fmt.Println(flag)
	// fmt.Println(data)
	// fmt.Println(string(data))

	// s1,err := b1.ReadString('\n')
	// fmt.Println(err)
	// fmt.Println(s1)

	// s1,err = b1.ReadString('\n')
	// fmt.Println(err)
	// fmt.Println(s1)

	// for{
	// 	s1,err := b1.ReadString('\n')
	// 	if err == io.EOF{
	// 		fmt.Println("读取完毕")
	// 		return
	// 	}
	// 	fmt.Println(s1)
	// }

	// data, err := b1.ReadBytes('\n')
	// fmt.Println(err)
	// fmt.Println(data)

	//Scaner
	// s2 := ""
	// fmt.Scanln(&s2)
	// fmt.Println(s2)

	// b2 := bufio.NewReader(os.Stdin)
	// s2, _ := b2.ReadString('\n')
	// fmt.Println(s2)


	

}
