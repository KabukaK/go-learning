package main

import (
	"demo/app/expense"
	"fmt"
)

func main() {

	var limit float64 = 1000
	expenses := []expense.Expense{
		{Amount: 500, Category: "Еда", Date: "2026-09-17", Paycard: true},
		{Amount: 1500, Category: "Машина", Date: "2026-09-15", Paycard: false},
		{Amount: 2000, Category: "Машина", Date: "2026-09-12", Paycard: false},
	}
	defer fmt.Println("Прогграмма окончена")

	expenses = append(expenses, expense.Expense{Amount: 1000, Category: "Прочее", Date: "2026-09-13", Paycard: false})
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
		go expense.ClassifyAsync(exp, limit, ch)
	}
	for i := 0; i < len(expenses); i++ {
		fmt.Println(<-ch)
	}

}
