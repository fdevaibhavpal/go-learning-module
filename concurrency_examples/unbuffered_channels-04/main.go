package main

import (
    "fmt"
    "time"
)

func sendData(ch chan int) {
    fmt.Println("Sending data to channel")
    ch <- 42 
}

func main() {
    ch := make(chan int) // create an Unbuffered channel

    go sendData(ch) 

    time.Sleep(1 * time.Second) 

    fmt.Println("Receiving data from channel")
    data := <-ch 
    fmt.Println("Data received from channel:", data)
}
