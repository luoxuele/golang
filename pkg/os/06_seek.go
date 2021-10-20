package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "test/heheda.txt"

	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//读写
	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4,io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))
}
