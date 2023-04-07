package main

import (
  "fmt"
  "time"
)



func say(s string) {
  for i := 0; i < 5; i++ {
    fmt.Println(s,"***",i)
  }
}


func main() {
  say("Sync") // 동기적 호출 

  // 비동기적 호출
  go say("Async1")
  go say("Async2")
  go say("Async3")

  // Wait for 3 seconds

  time.Sleep(time.Second * 3)
}
