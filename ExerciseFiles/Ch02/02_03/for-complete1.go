package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 20; i++ {
		//fmt.Println("i =", i)
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizz buzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
	fmt.Println("\n\nDone")
}
