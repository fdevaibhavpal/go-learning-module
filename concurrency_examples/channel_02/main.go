package main

import "fmt"

func fun(ch chan string) {
    ch <- "Hello,from channel!"
}

func main() {
    ch := make(chan string)
    go fun(ch)
    msg := <-ch
    fmt.Println(msg)
}
