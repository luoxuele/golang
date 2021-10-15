package main

import (
	"fmt"
	"net"
	"time"
)

/*
	net.Listen() //创建监听套接字
	Accept()  //阻塞等待用户连接
	Read()
	Write()
	Close()
*/

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println("net.Listen :", err)
		return
	}
	defer listener.Close()
	fmt.Println("监听socket ", listener.Addr().String())

	//阻塞监听
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept : ", err)
		return
	}

	defer conn.Close()
	fmt.Println("客户端连接成功： ", conn.LocalAddr(), conn.RemoteAddr())

	time.Sleep(time.Second * 99999)

	//读数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}

	fmt.Println("服务器读到数据： ", string(buf[:n]))

}
