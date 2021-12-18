package main

import "fmt"

// 매개변수타입, 리턴타입은 이름 뒤에 지정
func add1(x int, y int) int {
	return x + y
}

// 매개변수가 같은 타입일 경우 타입을 한번만 명시해줘도 된다.
func add2(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("add1의 결과 : ", add1(42, 11))
	fmt.Println("add2의 결과 : ", add2(42, 11))
}
