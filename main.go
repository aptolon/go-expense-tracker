package main

import (
	"expense-tracker/expense"
	"fmt"
)

func main() {
	s, err := expense.NewStorage("./str.json")
	if err != nil {
		fmt.Println("Error1:", err)
		return
	}
	e, err := s.AddExpense("Test", "2025-01-01", 200)
	if err != nil {
		fmt.Println("Error2:", err)
		return
	}
	e, err = s.AddExpense("Test1", "2025-01-01", 201)
	if err != nil {
		fmt.Println("Error2:", err)
		return
	}
	e, err = s.AddExpense("Test1", "", 201)
	if err != nil {
		fmt.Println("Error2:", err)
		return
	}
	fmt.Println(s.TotalSummary())
	fmt.Println(s.MonthlySummary(1))
	fmt.Println(s.GetAllExpenses())
	// e, err := s.UpdateExpense(1, "изменено", "", 200)
	// if err != nil {
	// 	fmt.Println("Error3:", err)
	// 	return
	// }
	// err = s.DeleteExpense(1)
	// if err != nil {
	// 	fmt.Println("Error4:", err)
	// 	return
	// }

	fmt.Println("Added:", e)
}
