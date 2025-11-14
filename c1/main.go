package main

import (
	"fmt"
)

func main() {

	fmt.Println("FizzBuzz Game: ")
	for num := 1; num <= 20; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("Fizz Buzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(num)
		}
	}

}
