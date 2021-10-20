package main

import (
	"fmt"
	"path"
)

func main(){
	str := path.Base("./")
	fmt.Println(str)
}  