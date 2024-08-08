package main

import "fmt"

const (
	Visa         = "Visa"
	Mastercard   = "Mastercard"
	AmericanExpress = "American Express"
)

func main() {
	cardTypes := []string{Visa, Mastercard, AmericanExpress}

	// Iterate over the slice and print each value
	for _, cardType := range cardTypes {
		fmt.Println(cardType)
	}
}
