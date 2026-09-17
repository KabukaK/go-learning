package main

import (
	"errors"
	"fmt"
)

type Expense struct {
	Amount   float64
	Category string
	Date     string
	Paycard  bool
}

func main() {

	var limit float64 = 1000
	expenses := []Expense{
		{Amount: 500, Category: "Еда", Date: "2026-09-17", Paycard: true},
		{Amount: 1500, Category: "Машина", Date: "2026-09-15", Paycard: false},
		{Amount: 2000, Category: "Машина", Date: "2026-09-12", Paycard: false},
	}
	defer fmt.Println("Прогграмма окончена")

	expenses = append(expenses, Expense{Amount: 1000, Category: "Прочее", Date: "2026-09-13", Paycard: false})
	for _, exp := range expenses {
		fmt.Println(exp.Classify(limit))
	}
	// fmt.Println(total(expenses))
	// fmt.Println(totalsByCategory(expenses))
	exp, err := NewExpense(100, "Машина")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Новая трата:", exp)
	}

}

func (e Expense) Classify(limit float64) (string, bool) {

	if e.Amount > limit {
		return "Крупная трата", true
	} else {
		return "Мелкая трата", false
	}

}

func total(expenses []Expense) float64 {
	var sum float64
	for _, exp := range expenses {
		sum += exp.Amount
	}
	return sum
}
func totalsByCategory(expenses []Expense) map[string]float64 {
	totals := make(map[string]float64)
	for _, exp := range expenses {
		totals[exp.Category] += exp.Amount
	}
	return totals
}

func NewExpense(amount float64, category string) (Expense, error) {
	if amount <= 0 || category == "" {
		return Expense{}, errors.New("сумма траты должна быть положительной")
	}
	return Expense{Amount: amount, Category: category}, nil
}
