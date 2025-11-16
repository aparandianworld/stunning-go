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

func safeAccess(values []int, index int) (int, error) {
	if index < 0 || index >= len(values) {
		return 0, errors.New("index out of bound")
	}

	return values[index], nil
}
