package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "82.156.32.144:9999")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}

	defer conn.Close()

	//写数据
	conn.Write([]byte("Are you ready?"))

	//接收数据
	buf := make([]byte,4096)
	n,err := conn.Read(buf)

	if err != nil{
		fmt.Println("conn.Read() err:　",err)
		return
	}

	fmt.Println(string(buf[:n]))

}
