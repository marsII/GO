package main

import (
	"fmt"
	"net"
)

func main() {
	A := "google.com:80"
	addr,err := net.ResolveTCPAddr("tcp",A)
	if err != nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(addr)
	}
}

//-------OUT-----------
172.217.23.238:80