package main

import (
    "fmt"
    "sync"
)

func workerFun(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    id++
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go workerFun(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers done")
}