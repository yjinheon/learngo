package main

import "fmt"

func Pointer() {
  var a int = 500
  var p *int  // init pointer

  p = &a

  fmt.Printf("p의 값 : %p\n",p)
  fmt.Printf("p의 메모리의 값 :%d\n",*p)

  *p =100
   
  fmt.Printf("a의 값 :%d\n",a)
}
