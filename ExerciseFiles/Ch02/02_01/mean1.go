// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {

	// x := 1.0
	// y := 2.0

	x, y := 1.0, 2.0

	fmt.Printf("x=%v, Type of %T\n", x, x)
	fmt.Printf("x=%v, Type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, Type of %T\n", mean, mean)
}
