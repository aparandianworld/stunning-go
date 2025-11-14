package main

import (
	"fmt"
)

func main() {

	a := 100.0
	b := 200.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("frac =", frac)
		fmt.Printf("a = %v is less than b = %v\n", a, b)
	} else {
		fmt.Println("frac =", frac)
		fmt.Printf("b = %v is less than a = %v\n", b, a)
	}

	num := 100

	switch num {
	case 10:
		fmt.Println("num = 10")
	case 20:
		fmt.Println("num = 20")
	case 30:
		fmt.Println("num = 30")
	default:
		fmt.Println("num is not 10 or 20 or 30")
	}

	number := 10

	// condition can also be in case
	switch {
	case number < 10:
		fmt.Println("number is less than 10")
	case number > 10:
		fmt.Println("number is greater than 10")
	default:
		fmt.Println("number is 10")
	}
}
