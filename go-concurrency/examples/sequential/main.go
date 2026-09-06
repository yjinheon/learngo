// sequential은 커피를 순차적으로 내리는 예제입니다.
package main

import (
	"fmt"
	"time"
)

func makeCoffee(name string) {
	fmt.Println("start: ", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Done : ", name)
}

func main() {
	makeCoffee("americano")
	makeCoffee("latte")
}
