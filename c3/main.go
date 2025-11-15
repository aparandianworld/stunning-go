package main

import (
	"fmt"
)

func main() {

	nums := []int{16, 23, 788, 4, 29, 54, 32, -4, 9, 12, 69, 420}
	// fmt.Println(nums)

	maximal := findMaximal(nums)
	fmt.Printf("Maximal value in slice nums %v is %d\n", nums, maximal)
}

func findMaximal(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	for _, num := range nums[1:] {
		if num > result {
			result = num
		}
	}
	return result
}
