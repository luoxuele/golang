package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	srcFile := "test/heheda.txt"
	destFile := "test2/xxx3.txt"
	// total, err := CopyFile1(srcFile, destFile)
	// total, err := CopyFile2(srcFile, destFile)
	total, err := CopyFile3(srcFile, destFile)
	fmt.Println(total, err)

}

func CopyFile3(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}

	return len(bs), nil
}

func CopyFile2(srcFile, destFile string) (int64, error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	dest, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer src.Close()
	defer dest.Close()
	return io.Copy(dest, src)
}

func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	bs := make([]byte, 1024, 1024)
	n := -1 //读取数据量
	total := 0

	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("file1.Read() err:", err)
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}

	return total, nil

}
