package main

import (
	"fmt"
	"time"
)

// You can do a lot with built-in types but sometimes you need to create your own
type Budget struct {
	Name         string
	ID           int
	amount       float64 // note the lowercase (not exported meaning not visible outside of this package)
	CreationDate time.Time
	ExiryDate    time.Time
}

func main() {

	budget := Budget{
		Name:         "Q1 Budget",
		ID:           1,
		amount:       1000.00,
		CreationDate: time.Now(),
		ExiryDate:    time.Now().Add(time.Hour * 24 * 30 * 3),
	}

	fmt.Println("Budget: ", budget)

	// if you want to see more info
	fmt.Printf("Budget: %#v\n", budget)

	// access fields using dot notation
	fmt.Printf("Budget Name: %s\n", budget.Name)
	fmt.Printf("Budget ID: %d\n", budget.ID)
	fmt.Printf("Budget Amount: %f\n", budget.amount)
	fmt.Printf("Budget Creation Date: %v\n", budget.CreationDate)
	fmt.Printf("Budget Expiry Date: %v\n", budget.ExiryDate)

	// zero value
	var defaultBudget Budget
	fmt.Println("Default Budget: ", defaultBudget)

}
