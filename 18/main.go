package main

import (
	"fmt"
)

func main() {
	ch1, ch2 := make(chan int, 5), make(chan int, 5)

	// sender 1
	go func() {
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
	}()

	// sender 2
	go func() {
		defer close(ch2)
		for j := 10; j > 5; j-- {
			ch2 <- j
		}
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Printf("channel 1 received: %d\n", v)
			} else {
				ch1 = nil
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Printf("channel 2 received: %d\n", v)
			} else {
				ch2 = nil
			}
		}
	}
}
