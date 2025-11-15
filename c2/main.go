package main

import (
	"fmt"
)

/* even-ended number is a number who's first and last digits are the same */
func main() {
	fmt.Println("How many even-ended numbers are there that are multiplication of 4 digit numbers?")
	counter := 0
	for i := 1000; i < 10000; i++ {
		for j := i; j < 10000; j++ { // don't count twice
			if i%2 == 0 && j%2 == 0 {
				product := i * j
				productString := fmt.Sprintf("%d", product)
				if productString[0] == productString[len(productString)-1] {
					counter++
				}
			}
		}
	}
	fmt.Println("Counter: ", counter)
}
