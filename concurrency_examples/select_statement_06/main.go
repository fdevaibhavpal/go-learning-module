package main

import (
    "fmt"
    "time"
)

func sendMsgToChannel(ch chan<- string, msg string, delay time.Duration) {
    time.Sleep(delay)
    ch <- msg
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
	d1 := 2 * time.Second
    d2 := 1 * time.Second

    go sendMsgToChannel(ch1, "message from ch1", d1)
    go sendMsgToChannel(ch2, "message from ch2", d2)

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout")
    }
}
