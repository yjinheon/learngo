package main

import "fmt"

func main() {
	var printVal string = "Hello Go"
	printMe(printVal)
}

func printMe(printVal string) {
	fmt.Println(printVal)
}
