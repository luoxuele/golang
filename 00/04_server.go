package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	//创建监听socket
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("net.Listen() err: ", err)
		return
	}
	defer listener.Close()

	//监听客户端连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept() err:", err)
			return

		}

		//具体完成服务器和客户端的数据通信
		go HandleConnect(conn)

	}

}

func HandleConnect(conn net.Conn) {
	defer conn.Close()

	//获取客户端的 Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, "客户端成功连接")

	//循环读客户端
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)

		if "exit\n" == string(buf[:n]) || "exit\n\r" == string(buf[:n]){
			fmt.Println("服务器接收到客户端退出请求，服务器断开连接",conn.RemoteAddr())
			return
		}

		if n == 0 {
			fmt.Println(conn.RemoteAddr(), " 客户端已关闭，服务端断开连接")
			return
		}

		if err != nil {
			fmt.Println("conn.Read() err: ", err)
			return
		}

		//小写转大写，回写给客户端
		str := strings.ToUpper(string(buf[:n]))
		conn.Write([]byte(str))

	}

}
