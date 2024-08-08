package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base: %d", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "Some Name",
	}
	fmt.Println("co={num: %v}")
}
