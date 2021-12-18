package main

import "fmt"
import "strings"

// go는 특이하게도   return 값을 2개로 받을 수 있다.

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {

	totalLengtht, upperName := lenAndUpper("nico")
	fmt.Println(totalLengtht, upperName)

}
