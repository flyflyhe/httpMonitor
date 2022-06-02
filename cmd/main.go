package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, _ := net.LookupIP("www.baidu.com")
	
	fmt.Println(addrs)
}
