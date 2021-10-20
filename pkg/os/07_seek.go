package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	/*
		断点续传：边写入边记录写入的数据总量
	*/

	srcFile := "test/HCIA-Datacom V1.0 培训教材.pdf"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:]
	fmt.Println(destFile)

	tempFile := destFile+".temp"
	fmt.Println(tempFile)

	src,err := os.Open(srcFile)
	HandleErr(err)
	dest,err := os.OpenFile(destFile,os.O_CREATE| os.O_WRONLY,os.ModePerm)
	HandleErr(err)
	temp,err := os.OpenFile(tempFile,os.O_CREATE| os.O_WRONLY,os.ModePerm)
	HandleErr(err)

	defer src.Close()
	defer dest.Close()
	defer temp.Close()

	temp.Seek(0,io.SeekStart)
	bs := make([]byte,100,100)
	n1,err := temp.Read(bs)

}

func HandleErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}