package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("www.snowfox.wang2")
	fmt.Println(err)
	fmt.Println(addr)
}
