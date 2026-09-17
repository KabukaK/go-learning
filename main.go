package main

import (
	"fmt"
)

func main() {

	// var amount float64 = 500
	// var category string = "Еда"
	// var datespend string = "2026-09-17"
	// var paycard bool = true
	var limit float64 = 1000
	e := Expense{Amount: 500, Category: "Еда", Date: "2026-09-17", Paycard: true}
	defer fmt.Println("Прогграмма окончена")
	// fmt.Println(classify(amount, limit))
	for i := 1; i <= 7; i++ {
		fmt.Println(i)
	}
	// fmt.Println(amount, category, datespend, paycard)
	fmt.Println(total(500, 100, 1000))
	fmt.Println(e.Classify(limit))
}

func (e Expense) Classify(limit float64) (string, bool) {

	if e.Amount > limit {
		return "Крупная трата", true
	} else {
		return "Мелкая трата", false
	}

}

func total(amounts ...float64) float64 {
	var sum float64
	for _, amount := range amounts {
		sum += amount
	}
	return sum
}

type Expense struct {
	Amount   float64
	Category string
	Date     string
	Paycard  bool
}
