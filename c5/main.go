package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "http://www.google.com/"

	cType, err := contentType(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	fmt.Printf("URL: %s\nContent-Type: %s\n", url, cType)

}

func contentType(url string) (string, error) {

	// aquire resource, check for errors and release resource w/ defer pattern
	response, err := http.Get(url)

	if response == nil {
		return "", fmt.Errorf("response is nil")
	}

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	return response.Header.Get("Content-Type"), nil
}
