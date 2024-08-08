package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func vals2() (int, int, int) {
	return 3, 7, 9
}

func main() {
	a, b := vals() // 여러개의 리턴값을 받을 수 있음
	c, d, e := vals2()
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
