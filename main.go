package main

import (
	"expense-tracker/expense"
	"fmt"
)

func main() {
	s, err := expense.NewStorage("/Users/admin/go_learn/expense-tracker/str.json")
	if err != nil {
		fmt.Println("Error1:", err)
		return
	}
	e, err := s.AddExpense("Test", "2024-01-01", 200)
	if err != nil {
		fmt.Println("Error2:", err)
		return
	}

	fmt.Println("Added:", e)
}
