package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	
	addrs, err := net.LookupHost("luo.snowfox.wang")
	// addrs, err := net.LookupTXT("www.snowfox.wang")
	// addrs, err := net.LookupAddr("20.205.243.166")

	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(addrs)
}
