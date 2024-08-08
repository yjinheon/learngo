package learn

import "fmt"

func plus(a int, b int) int { // 두개의 int를 받아서 int를 리턴하는 함수
	return a + b
}

func dplus(a, b, c int) int { // 세개의 int를 받아서 int를 리턴하는 함수
	return a + b + c
}

func LearnFunc() {
	var line string = "##### 12_functions.go #####"
	fmt.Println(line)
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = dplus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
