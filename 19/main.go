package main

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AdURL string
	Price int
}

var defaultBid = Bid{
	AdURL: "http://example.default.bid.com",
	Price: 50,
}

func bestBid(url string) Bid {
	time.Sleep(25 * time.Millisecond)
	return Bid{
		AdURL: url,
		Price: 100,
	}
}

func processBid(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1) // buffered channel

	go func() {
		ch <- bestBid(url)
	}()

	select {

	case bid := <-ch: // bid received in allocated time
		return bid

	case <-ctx.Done(): // timeout
		return defaultBid

	}
}

func main() {
	timeout := 100 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	url := "http://example.bid.cat.com"
	bid := processBid(ctx, url)
	fmt.Println("enough time", bid)

	// not enough time
	timeout = 10 * time.Millisecond
	ctx, cancel = context.WithTimeout(context.Background(), timeout)
	defer cancel()

	url = "http://example.bid.dog.com"
	bid = processBid(ctx, url)
	fmt.Println("not enough time", bid)

}
