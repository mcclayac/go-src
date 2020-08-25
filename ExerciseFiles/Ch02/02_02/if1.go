// Example of "it" statement
package main

import (
	"fmt"
)

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is Big")
	}

	if x > 100 {
		fmt.Println("x is very big")
	} else {
		fmt.Println("x is not that big")
	}
	fmt.Printf("The value of x is %v\n\n", x)

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	if x > 20 || x > 30 {
		fmt.Println("x is out of range")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Printf("a is more than hal of b, a=%.2f b=%.2f frac=%.2f\n\n",
			a, b, frac)
	}
}
