package main

import (
	"os"
	"fmt"
)

func main(){
	f,e := os.Stdin.Stat();
	_ = f
	if e != nil {
		panic(e)
	}
	if  os.ModeNamedPipe != 0 {
		fmt.Print("Hi pipe!")
	}else{
		fmt.Print("no anything")
	}
}