package main

import "fmt"

// intseq는 함수를 반환한다

func intSeq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main() {

  nextInt := intSeq()
  
  i := 1 // i 초기화 
  for i <= 20 {
      fmt.Println(nextInt())
      i = i + 1
  }

  newInts := intSeq()
  fmt.Println(newInts())

}
