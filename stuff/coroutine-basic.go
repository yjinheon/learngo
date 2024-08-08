package main

import (
    "fmt"
    "time"
)

func myCoroutine(ch chan string) {
    fmt.Println("Coroutine started")
    time.Sleep(time.Second)  // Pause for one second
    ch <- "Coroutine resumed"
    time.Sleep(2 * time.Second)  // Pause for two seconds
    ch <- "Coroutine finished"
}

func main() {
    fmt.Println("Main started")
    ch := make(chan string) // make : channel 생성
    go myCoroutine(ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println("Main finished")
}
