package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "test/heheda.txt"
	file, err := os.OpenFile(filename,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bs := []byte{65, 66, 67, 68, 69, 70}
	n, err := file.Write(bs)
	HandleErr(err)
	fmt.Println(n)

	file.WriteString("嘿嘿嘿")

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
