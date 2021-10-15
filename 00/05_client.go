package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	//建立通信套接字
	conn, err := net.Dial("tcp", "82.156.32.144:9999")
	if err != nil {
		fmt.Println("net.Dial() err: ", err)
		return
	}

	defer conn.Close()

	//	获取键盘输入（stdin），并将输入发送给服务器
	go func() {
		str := make([]byte, 4096)

		for {
			n, err := os.Stdin.Read(str)
			if err != nil{
				fmt.Println("os.Stdin.Read() err: ",err)
				continue
			}

			//写给服务器
			conn.Write(str[:n])


		}
	}()

	// 回显服务器回发的大写数据
	buf := make([]byte, 4096)
	for {
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read() err: ",err)
			return
		}

		fmt.Println(string(buf[:n]))

	}

}
