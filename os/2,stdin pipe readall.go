package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main(){
	f,e := os.Stdin.Stat();
	_ = f
	if e != nil {
		panic(e)
	}
	data , e := ioutil.ReadAll(os.Stdin)
	fmt.Printf("input is %v", string(data) )
}

/**
1.
go run FILE < another file
2.
echo TEXT | go run FILE
**/