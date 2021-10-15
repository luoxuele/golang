package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("net.Listen: ", err)
		return
	}

	defer listen.Close()
	//fmt.Println(listen.Addr())
	listen.Accept()

	time.Sleep(time.Second * 9999)

	//coon , err := listen.Accept()
	//fmt.Println("helo")
}
