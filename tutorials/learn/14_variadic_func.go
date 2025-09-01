package learn

import "fmt"

// variadic function은 인자를 여러개 받을 수 있는 함수를 의미한다.

func sum(nums ...int) { // int 여러개를 인자로 받음

	fmt.Println(nums, " ")
	total := 0 // `:=` is a declaration , whreas `=` is an assignment
	// the difference is that `:=` is used to declare a new variable

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func VariadicFunc() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
