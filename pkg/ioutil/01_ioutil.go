package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// fileName := "../bufio/test.txt"
	// data, err := ioutil.ReadFile(fileName)
	// fmt.Println(err)
	// fmt.Println(string(data))

	// fileName := "test.txt"
	// s1 := "aaahelloworld面朝大海，春暖花开"
	// err := ioutil.WriteFile(fileName,[]byte(s1),os.ModePerm)
	// fmt.Println(err)

	// s2 := "heheda呵呵哒0"
	// r1 := strings.NewReader(s2)
	// data, err := ioutil.ReadAll(r1)
	// fmt.Println(err)
	// fmt.Println(data)

	dirName := "/root"
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(fileInfos)
	for i:=0;i<len(fileInfos);i++{
		// fmt.Printf("%T\n",fileInfos[i])
		fmt.Printf("%d %s %t\n",i,fileInfos[i].Name(),fileInfos[i].IsDir())
	}

}
