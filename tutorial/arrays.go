package main

import "fmt"

func main() {
	var a [5]int // 크기5의 비어있는 배열 생성
	fmt.Println("emp: ",a)

	a[4] = 100 // 5번째에 100할당

	fmt.Println("set: ",a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1,2,3,4,5} // 배열에 여러 값을 한번에 넣을 때는 중괄호 사용
	fmt.Println("dcl: " ,b)

	var twoD [2][3]int // 2차원 배열 생성
	for i:=0; i<2; i++{ // comma 가 아니라 ;로 구분함
		for j := 0; j < 3; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
