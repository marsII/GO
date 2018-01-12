package main


import (
	"fmt"
	"time"
)

func main() {
    naturals := make(chan int)  //channel
    squares := make(chan int)

    // Counter
    go func() {
        for x := 0; x < 10 ; x++ {
            naturals <- x
			fmt.Printf("i set x = %d\n",x)
			time.Sleep(1000 * time.Millisecond)
        }
		close(naturals)
    }()

    // Squarer
    go func() {
        for x := range naturals {
            squares <- x
			fmt.Printf("i read x = %d\n",x)
        }
		close(squares)
    }()

    // Printer (in main goroutine)
    for x := range squares{
        fmt.Println(x)
    }
}



/**
i set x = 0
i read x = 0
0
i set x = 1
i read x = 1
1
i set x = 2
i read x = 2
2
i set x = 3
i read x = 3
3
i set x = 4
i read x = 4
4
i set x = 5
i read x = 5
5
i set x = 6
i read x = 6
6
i set x = 7
i read x = 7
7
i set x = 8
i read x = 8
8
i set x = 9
i read x = 9
9

**/