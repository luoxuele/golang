package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("test/readme.txt")
	if err != nil {
		fmt.Println("os.Open() err: ", err)
		return
	}

	defer file.Close()

	bs := make([]byte,4,4)
	// n,err := file.Read(bs)
	// fmt.Println(err)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// n,err = file.Read(bs)
	// fmt.Println(err)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// n,err = file.Read(bs)
	// fmt.Println(err)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	// n,err = file.Read(bs)
	// fmt.Println(err)
	// fmt.Println(n)
	// fmt.Println(string(bs))

	n := -1
	for{
		n, err = file.Read(bs)
		if n==0 || err == io.EOF{
			fmt.Println("read over")
			break
		}

		fmt.Println(string(bs[:n]))
	}
}
