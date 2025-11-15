package main

import (
	"fmt"
	"log"
	"math"
)

type Square struct {
	X      int
	Y      int
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {

	if length <= 0 {
		return nil, fmt.Errorf("length must be greater than 0")
	}

	square := Square{
		X:      x,
		Y:      y,
		Length: length,
	}

	return &square, nil
}

func (s *Square) Move(dx int, dy int) { // receiver pointer of type Square
	s.X += dx
	s.Y += dy
}

func (s *Square) Area() float64 { // receiver value of type Square
	return math.Pow(float64(s.Length), 2)
}

func main() {

	square, err := NewSquare(1, 1, 10)
	fmt.Printf("Original Square: %#v\n", square)
	if err != nil {
		log.Fatalf("Error creating square: %v", err)
	}

	area := square.Area()
	fmt.Println("Area:", area)

	square.Move(2, 3)
	fmt.Printf("Moved Square: %#v\n", square)

}
