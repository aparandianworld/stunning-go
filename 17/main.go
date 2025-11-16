package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channels are used for communication between goroutines. They are one-direction pipes with specific type")
	fmt.Println("receiver < [channel] < sender")

	/*
		// send a value to channel
		ch := make(chan int)
		ch <- 19

		// receive a value from channel (panic)
		receiver := <-ch
		fmt.Printf("receiver: %d\n", receiver)
	*/

	ch := make(chan int)
	go func() {
		ch <- 19
	}()

	receiver := <-ch
	fmt.Printf("received: %d\n", receiver)

	fmt.Println("---")

	const counter = 10
	for i := 1; i <= counter; i++ {

		go func() {
			fmt.Printf("sending: %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}()
	}

	for j := 1; j <= counter; j++ {
		receiver := <-ch
		fmt.Printf("received: %d\n", receiver)
	}

	fmt.Println("---")

	go func() {
		for i := 1; i < counter; i++ {
			fmt.Printf("sending: %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}

		close(ch)
	}()

	for receiver := range ch {
		fmt.Printf("received: %d\n", receiver)
	}

	fmt.Println("---")

	bch := make(chan int, 3) // buffered channel - in this case I can store 3 value
	bch <- 100
	bch <- 200
	bch <- 300

	for i := 0; i < 3; i++ {
		bufferedReceiver := <-bch
		fmt.Println("buffered channel received: ", bufferedReceiver)
	}

}
