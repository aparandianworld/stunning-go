package main

import "fmt"

// generic type - in this case fields can be int, int64, float64 or string
type Ordered interface {
	int | int64 | float64 | string
}

func min[T Ordered](values []T) (T, error) { // min function of generic type T

	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("Empty slice")
	}

	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return min, nil
}

func main() {

	// test all
	minInt, err := min([]int{9, 10, -1, 22})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Min int: %d\n", minInt)

	minString, err := min([]string{"apple", "banana", "orange"})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Min string: %s\n", minString)

	minFloat, err := min([]float64{9.1, 10.2, -1.3, 22.4})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Min float: %f\n", minFloat)

	// test empty
	emptySlice, err := min([]int{})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("Min empty: %d\n", emptySlice)
}
