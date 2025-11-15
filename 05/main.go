package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice: ")

	loons := []string{"bugs", "daffy", "donald"}

	fmt.Println(loons)
	fmt.Println("Length:", len(loons))
	fmt.Println("Capacity:", cap(loons))
	fmt.Println("---")
	fmt.Println("loons[0]]: ", loons[0])
	fmt.Println("loons[1]: ", loons[1])
	fmt.Println("loons[2]: ", loons[2])
	fmt.Println("---")
	fmt.Println("loons[1:2]: ", loons[1:2])
	fmt.Println("---")

	// loops
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}
	fmt.Println("---")

	for index, name := range loons {
		fmt.Println("index: ", index, "name: ", name)
	}
	fmt.Println("---")

	fmt.Println("---")
	loons = append(loons, "elmer")
	fmt.Println(loons)
	fmt.Println("Length:", len(loons))
	fmt.Println("Capacity:", cap(loons))
	fmt.Println("---")

	for _, name := range loons {
		fmt.Println("name: ", name)
	}
	fmt.Println("---")

}
