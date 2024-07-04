package main

import (
    "fmt"
   "time"
)

func hello() {
    fmt.Println("Hello, World")
}

func main() {
    go hello() 
    time.Sleep(1 * time.Second) 
    fmt.Println("Learning go language")
}
