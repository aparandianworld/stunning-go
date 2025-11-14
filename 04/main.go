package main

import (
	"fmt"
)

func main() {
	book := "The Go Programming Language"

	fmt.Println("Book: ", book)
	fmt.Println("Length of Book: ", len(book))

	fmt.Println("Book[0] = ", book[0])
	fmt.Println("Book[0] = ", string(book[0]))

	for _, ch := range book {
		fmt.Printf("ch = %v, type: %T\n", ch, ch)
	}

	// slice
	fmt.Println("Book[0:4] = ", book[0:4])
	fmt.Println("Book[:5] = ", book[:5])
	fmt.Println("Book[4:] = ", book[4:])

	// raw multiline string
	quote := `Let me tell you a secret: no one does when they begin. Ideas don't come out fully formed. They only become clear as you work on them. You just have to get started...`
	fmt.Println("Mark Zuckerberg's quote: ", quote)

}
