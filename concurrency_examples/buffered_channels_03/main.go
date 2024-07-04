package main

import "fmt"

func bufChan(ch chan int) {
	ch <- 4
}

func main() {
    ch := make(chan int, 3)   // buffered channel with capacity 3

    ch <- 1
    ch <- 2
    ch <- 3

    fmt.Println(<-ch) 
    fmt.Println(<-ch) 

    go bufChan(ch) 

    fmt.Println(<-ch) 
    fmt.Println(<-ch) 
}
