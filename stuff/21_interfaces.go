package main

import (
  "fmt"
  "math"
)

// interfaces are named collections of method signature
// go에서 interface는 method들의 집합체이다.
// interface는 특정 값이 가지고 있기를 기대하는 메서드 집합이다.
// 기본적으로 어떤 값이 이 특정 타입을 갖는지 관심이 없는 경우에 사용한다. 

type geometry interface {
     area() float64
     perim() float64
}

type rect struct {
     width, height float64
}

type circle struct {
     radius float64
}

func (r rect) area() float64 {
      return r.width * r.height
}

func (r rect) perim() float64 {
      return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
      return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
      return 2*math.Pi * c.radius
}

func measure(g geometry) {
     fmt.Println(g)
     fmt.Println(g.area())
     fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3 , height: 4}
    c := circle{radius:5}

    measure(r)
    measure(c)
}
