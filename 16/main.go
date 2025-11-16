package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func getContentType(url string) (contentType string, err error) {

	if url == "" {
		return "", errors.Wrap(err, "url is empty")
	}

	resp, err := http.Get(url)

	if err != nil {
		return "", errors.Wrap(err, "failed to get url")
	}

	defer resp.Body.Close()

	return resp.Header.Get("Content-Type"), nil

}

func fetchURLs(urls []string) (err error) {

	if len(urls) == 0 {
		return errors.Wrap(err, "urls slice is empty")
	}

	for _, url := range urls {
		cType, err := getContentType(url)
		if err != nil {
			return errors.Wrap(err, "failed to get content type")
		}
		fmt.Println("URL: ", url, "ContentType: ", cType)
	}

	return nil
}

func fetchURLsConcurrently(urls []string) (err error) {

	if len(urls) == 0 {
		return errors.Wrap(err, "urls slice is empty")
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			cType, err := getContentType(url)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			fmt.Println("URL: ", url, "ContentType: ", cType)
			wg.Done()
		}(url)
	}

	wg.Wait()
	return nil
}

func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.linkedin.com",
	}

	startTime := time.Now()
	fetchURLs(urls)
	endTime := time.Now()
	fmt.Println("Time duration (serial): ", endTime.Sub(startTime))

	fmt.Println("---\n")

	startTime = time.Now()
	fetchURLsConcurrently(urls)
	endTime = time.Now()
	fmt.Println("Time duration (concurrently): ", endTime.Sub(startTime))

}
