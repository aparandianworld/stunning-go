package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Request struct {
	User   string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

var data = `
[
  {
    "user": "john doe",
    "type": "deposit",
    "amount": 100
  },
  {
    "user": "jane doe",
    "type": "withdraw",
    "amount": 500
  }
]
`

func main() {

	reader := strings.NewReader(data)
	decoder := json.NewDecoder(reader)

	var requests []Request

	if err := decoder.Decode(&requests); err != nil {
		log.Fatalf("Error decoding requests: %s", err)
	}

	for index, req := range requests {
		fmt.Printf("Decoded request %d: %+v\n", index, req)
	}
}
