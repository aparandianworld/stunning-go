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

func main() {

	budget := Budget{
		Name:           "Q2 Budget",
		ID:             2,
		Amount:         2000.00,
		CreationDate:   time.Now(),
		ExpirationDate: time.Now().Add(time.Hour * 24 * 30 * 3),
	}

	fmt.Printf("Detailed Budget: %#v\n", budget)

	budget.updateBudgetAmount(1250.90)
	fmt.Println("Updated Budget: ", budget)
	fmt.Println("Time Remaining on Budget: ", budget.timeRemaining())
}

func (b *Budget) updateBudgetAmount(amount float64) { // (b *Budget) b is a pointer receiver of type Budget
	b.Amount += amount
}

func (b Budget) timeRemaining() time.Duration { // (b Budget) b is a value receiver of type Budget
	return b.ExpirationDate.Sub(b.CreationDate)
}
