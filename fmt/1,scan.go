package main

import "fmt"

func main() {
	var input string
	for{
		input = ""
		fmt.Print("Enter text: ")
		
		fmt.Scanln(&input)
		
		fmt.Print(input)
		fmt.Printf("\n")
		
		if input == "bye" {
			break
		}
	}
}

/***
Scan scans text read from standard input,
 storing successive space-separated values into successive arguments.
 Newlines count as space.
---
Scanln is similar to Scan,
 but stops scanning at a newline and after the final item
 there must be a newline or EOF.

***/