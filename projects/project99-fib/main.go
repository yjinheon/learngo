package main

import (
	"fmt"
	"project99-fib/fib"
)

func main() {
	sequence := fib.FibonacciSequence(10)
	fmt.Println("FibonacciSequence of first 10 numbers:")
	fmt.Println(sequence)
}
