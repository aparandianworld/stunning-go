package main

import (
	"fmt"
	"io"
	"os"
)

/*
Implements io.Writer and turns everything to uppercase

	type Writer interface {
		Write(p []byte) (n int, err error)
	}
*/
type Capitalizer struct {
	Writer io.Writer
}

func (c *Capitalizer) Write(p []byte) (n int, err error) {

	asciiDiff := byte('a' - 'A') // 32
	out := make([]byte, len(p))

	for i, ch := range p {
		if ch >= 'a' && ch <= 'z' {
			ch -= asciiDiff // subtract 32 to convert lowercase to uppercase
		}
		out[i] = ch
	}

	n, err = c.Writer.Write(out)
	return n, err
}

func main() {

	s := &Capitalizer{Writer: os.Stdout}
	fmt.Fprintln(s, "This is an implementation of io.Writer interface and turns everything to uppercase")

}
