package main

import (
	"fmt"
)

func main() {
	fmt.Println("Garbage collection and defer...")

	worker()

	fmt.Println("Main thread...")
}

/*
Acquire a resource, check for error and release with defer is a common pattern in Go.
Note: release is in reverse order of acquire.
*/
func worker() {
	fmt.Println("Worker thread begins...")

	ra, err := aquire("Resource A")
	if err != nil {
		fmt.Println("Error acquiring resource A:", err)
		return
	}

	defer release(ra)

	rb, err := aquire("Resource B")
	if err != nil {
		fmt.Println("Error acquiring resource B:", err)
		return
	}

	defer release(rb)

	fmt.Println("Worker thread ends...")
}

func aquire(name string) (string, error) {
	fmt.Println("Acquiring resource: ", name)
	return name, nil
}

func release(name string) {
	fmt.Println("Releasing resource: ", name)
}
