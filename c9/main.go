package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func getContentType(url string, out chan string) {

	if url == "" {
		fmt.Println("Error: empty URL")
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s", url, err)
		return
	}

	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	out <- fmt.Sprintf("%s -> contentType: %s", url, contentType)
}

func fetchURLs(urls []string) (err error) {
	if len(urls) == 0 {
		return errors.Wrap(err, "empty urls")
	}

	rch := make(chan string)

	for _, url := range urls {
		go getContentType(url, rch)
	}

	for range urls {
		out := <-rch
		fmt.Println(out)
	}

	close(rch)

	return nil
}

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.linkedin.com",
	}

	fetchURLs(urls)
}
