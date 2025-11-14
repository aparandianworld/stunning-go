package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("---")

	for j := 0; j < 10; j++ {
		if j == 5 {
			break
		}
		fmt.Println("j = ", j)
	}

	fmt.Println("---")

	for k := 0; k < 10; k++ {
		if k == 5 {
			continue
		}
		fmt.Println("k = ", k)
	}

	fmt.Println("---")

	var m int
	m = 1

	for m < 10 { // similar to while loop in other languages
		fmt.Println("m = ", m)
		m++
	}

	fmt.Println("---")

	p := 1

	for { // infinite loop
		fmt.Println("p = ", p)
		p++
		if p == 7 {
			break
		}
	}

}
