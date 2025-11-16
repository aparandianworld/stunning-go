package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func killService(pidFile string) (err error) {

	if pidFile == "" {
		return errors.Wrap(err, "Error: pidFile is empty")
	}

	// common pattern again: open a resource, error handling and close w/ defer
	fh, err := os.Open(pidFile)
	if err != nil {
		return errors.Wrap(err, "Error: failed to open pidFile")
	}
	defer fh.Close()

	cb, err := io.ReadAll(fh)
	if err != nil {
		return errors.Wrap(err, "Error: failed to read pidFile")
	}

	pid, err := strconv.Atoi(string(cb))
	if err != nil {
		return errors.Wrap(err, "Error: failed to parse content from string to int")
	}
	// simulate a kill
	fmt.Printf("Killing service with pid = %d\n", pid)

	return nil
}

func main() {
	pf := "./server.pid"
	if err := killService(pf); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
