package main

import (
	"fmt"
	"math"
)

type Circle struct {
	X      int
	Y      int
	Radius int
}

type Square struct {
	X      int
	Y      int
	Length int
}

type Shape interface {
	Area() float64
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

func (s *Square) Area() float64 {
	return math.Pow(float64(s.Length), 2)
}

func main() {
	circle := Circle{
		X: 1, Y: 1, Radius: 4,
	}
	fmt.Printf("Circle: %#v\n", circle)
	fmt.Printf("Circle area: %f\n", circle.Area())

	square := Square{
		X: 1, Y: 1, Length: 4,
	}
	fmt.Printf("Square: %#v\n", square)
	fmt.Printf("Square area: %f\n", square.Area())

	shapes := []Shape{&circle, &square}
	result := 0.0
	for _, shape := range shapes {
		result += shape.Area()
	}
	fmt.Printf("Total area: %f\n", result)
}
