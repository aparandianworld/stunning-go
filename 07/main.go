package main

import (
	"fmt"
	"math"
)

func main() {

	r1, err := squareRoot(9)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(r1)
	}

	r2, err := squareRoot(-4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(r2)
	}

}

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0.0, fmt.Errorf("square root of negative number %v", x)
	}
	return math.Sqrt(x), nil
}
