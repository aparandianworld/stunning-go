package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {

	/*
		values := []int{1, 2, 3, 4, 5}

		fmt.Println("values: ", values)
		fmt.Println("length: ", len(values))

		_, err := fmt.Println(values[10])
		if err != nil {
			fmt.Println(errors.Wrap(err, "out of bound access"))
		} // panic: out of bound access

		panic("out of bound access...")

	*/

	values := []int{1, 2, 3, 4, 5}
	fmt.Println("values: ", values)
	fmt.Println("length: ", len(values))

	index := -3
	v, err := safeAccess(values, index)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error: "))
	}
	fmt.Printf("value at index %d is %d\n", index, v)

}

func safeAccess(values []int, index int) (n int, err error) {

	defer func() {
		if r := recover(); r != nil { // recover() is built-in function to capture error from panic
			err = fmt.Errorf("%v", r) // convert panic to error instead of crashing the program
		}
	}()

	return values[index], nil
}
