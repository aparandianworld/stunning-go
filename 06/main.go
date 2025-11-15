package main

import (
	"fmt"
)

func main() {

	fmt.Println("pass by value...")
	a := 2
	doubleValue(a)
	fmt.Println(a)

	fmt.Println("pass by reference...")
	slice := []int{1, 2, 10, -2, 6}
	doubleAt(slice, 2)
	fmt.Println(slice)

	fmt.Println("pass by reference...")
	b := 4
	doubleReference(&b)
	fmt.Println(b)
	fmt.Println(&b)
}

func doubleValue(a int) {
	a *= 2
}

func doubleAt(slice []int, index int) {
	slice[index] *= 2
}

func doubleReference(b *int) {
	*b *= 2
}
