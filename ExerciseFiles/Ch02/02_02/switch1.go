// Example of switch statement
package main

import (
	"fmt"
)

func main() {
	x := 2

	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("many")
	}

	switch {
	case x > 100:
		fmt.Println("many")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small")
	}
}
