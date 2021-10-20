package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir("/root/node-v14.17.3-linux-x64.tar.xz")
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		fmt.Println(i,file.Name())
	}
}
