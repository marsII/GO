package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"flag"
)

func main(){
	flag.Parse()
	
	var data []byte
	var e error
	
	_=e
	
	switch flag.NArg(){
		case 0:
			data , e = ioutil.ReadAll(os.Stdin)
			break
		case 1:
			data , e = ioutil.ReadFile(flag.Arg(0))
			break
	}
	
	fmt.Printf("input is %v", string(data) )
}

/**
1.
go run FILE < another file
2.
echo TEXT | go run FILE
3.
go run FILE txtfile.TXT
**/