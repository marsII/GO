package main

import (
	"fmt"
	"net"
	//"os"
	"bufio"
	//"io/ioutil"
)

func main(){
	var addr string
	addr = "127.0.0.1:4545"
	A , _ := net.ResolveTCPAddr("tcp",addr)
	con , _ := net.DialTCP("tcp",nil,A)
	
	r := bufio.NewReader(con)
	for {
		text , err := r.ReadString('\n')
		if err != nil 	{	break			}
		if len(text) > 0 {	fmt.Print(text) }
	}
}

/**
SERVER: socat tcp-l:4545,reuseaddr - OR nc -l -p 4545
CLINET: go run THIS FILE
and type in server ..
clinet print it
**/