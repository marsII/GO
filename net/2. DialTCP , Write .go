package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	A := "127.0.0.1:4545"
	addr,err := net.ResolveTCPAddr("tcp",A)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	con ,err := net.DialTCP("tcp",nil,addr)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	
	_ , err = con.Write([]byte("HI!"))
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}
	con.Close()
}

