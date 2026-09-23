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
	// for _, exp := range expenses {
	// 	fmt.Println(exp.Classify(limit))
	// }
	// fmt.Println(total(expenses))
	// fmt.Println(totalsByCategory(expenses))
	// exp, err := NewExpense(100, "Машина")
	// if err != nil {
	// 	fmt.Println("Ошибка:", err)
	// } else {
	// 	fmt.Println("Новая трата:", exp)
	// }

	// fmt.Println(filterByCategory(expenses, "Машина"))
	// fmt.Println(largestExpense(expenses))
	// fmt.Println(countByCategory(expenses))
	// expenses, err := addExpnense(expenses, 1000, "Еда")
	// if err != nil {
	// 	fmt.Println("Ошибка:", err)
	// }

	ch := make(chan string)
	for _, exp := range expenses {
		go classifyAsync(exp, limit, ch)
	}
	for i := 0; i < len(expenses); i++ {
		fmt.Println(<-ch)
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

func (e Expense) String() string {
	return fmt.Sprintf("%s: %.2f (%s)", e.Category, e.Amount, e.Date)
}

func filterByCategory(expenses []Expense, category string) []Expense {
	expense := []Expense{}
	for _, exp := range expenses {
		if exp.Category == category {
			expense = append(expense, Expense{Amount: exp.Amount, Category: exp.Category, Date: exp.Date, Paycard: exp.Paycard})
		}
	}
	return expense
}

func largestExpense(expenses []Expense) Expense {
	var maxAmount float64
	var ex Expense
	for _, exp := range expenses {
		if exp.Amount > maxAmount {
			maxAmount = exp.Amount
			ex = Expense{Amount: exp.Amount, Category: exp.Category, Date: exp.Date, Paycard: exp.Paycard}
		}

	}
	return ex
}

func countByCategory(expenses []Expense) map[string]int {
	counts := make(map[string]int)
	for _, exp := range expenses {
		counts[exp.Category] += 1
	}
	return counts
}

func addExpnense(expense []Expense, amount float64, category string) ([]Expense, error) {
	exp, err := NewExpense(amount, category)
	if err != nil {
		return expense, err
	} else {
		expense = append(expense, exp)
		return expense, nil
	}
}

func classifyAsync(exp Expense, limit float64, ch chan string) {
	message, _ := exp.Classify(1000)
	category := exp.Category
	amount := exp.Amount
	s := fmt.Sprintf("%s: %.2f - %s", category, amount, message)
	ch <- s
}
