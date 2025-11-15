package main

import (
	"fmt"
	"time"
)

type Budget struct {
	Name           string
	ID             int
	Amount         float64
	CreationDate   time.Time
	ExpirationDate time.Time
}

func NewBudget(name string, id int, amount float64, creationTime time.Time, expirationTime time.Time) (*Budget, error) {

	if name == "" {
		return nil, fmt.Errorf("Error name cannot be empty")
	}

	if id <= 0 {
		return nil, fmt.Errorf("Error id must be greater than 0")
	}

	if amount <= 0 {
		return nil, fmt.Errorf("Error amount must be greater than 0")
	}

	if expirationTime.Before(creationTime) {
		return nil, fmt.Errorf("Error expiration time must be after creation time")
	}

	budget := Budget{
		Name:           name,
		ID:             id,
		Amount:         amount,
		CreationDate:   creationTime,
		ExpirationDate: expirationTime,
	}

	return &budget, nil

}

func main() {

	budget := Budget{
		Name:           "Q3 Budget",
		ID:             2,
		Amount:         2200.00,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().Add(time.Hour * 24 * 30 * 3),
	}

	fmt.Printf("Detailed Budget: %#v\n", budget)

	newBudget, err := NewBudget("Q3 Budget v2", 3, 5500.00, time.Now(), time.Now().Add(time.Hour*24*30*3))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Detailed Budget v2: %#v\n", newBudget)

}
