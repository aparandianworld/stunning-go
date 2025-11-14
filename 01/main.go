package main

import (
	"fmt"
)

func main() {
	var x int
	var y int
	var a float64
	var b float64

	x = 10
	y = 20
	a = 16.2
	b = 4.1

	// Declare and assign in a single step (go will infer the type)
	// c, d := 20.2, 5.1 // multiple assignment (same as below)
	c := 20.2
	d := 5.1

	fmt.Printf("x = %v, type of x = %T\n", x, x)
	fmt.Printf("x = %v, type of x = %T\n", x, x)
	fmt.Printf("a = %v, type of a = %T\n", a, a)
	fmt.Printf("b = %v, type of b = %T\n", b, b)
	fmt.Printf("c = %v, type of c = %T\n", c, c)
	fmt.Printf("d = %v, type of d = %T\n", d, d)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of result: %T\n", mean, mean)
}
