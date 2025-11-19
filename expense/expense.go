package expense

import "time"

type Expense struct {
	Id          int
	Description string
	Date        time.Time
	Amount      float64
}

func NewExpense(id int, description string, dateString string, amount float64) (Expense, error) {
	var date time.Time
	if dateString != "" {
		date, _ = time.Parse("2006-01-02", dateString)
	} else {
		date = time.Now()
	}
	return Expense{
		Id:          id,
		Description: description,
		Date:        date,
		Amount:      amount,
	}, nil
}

// func TotalSummary(expenses map[int]Expense) float64 {
// 	var summ float64
// 	for _, e := range expenses {
// 		summ += e.Amount
// 	}
// 	return summ
// }

// func MonthlySummary(expenses map[int]Expense, month int) float64 {
// 	var summ float64
// 	year := time.Now().Year()
// 	for _, e := range expenses {
// 		if int(e.Date.Month()) == month && year == e.Date.Year() {
// 			summ += e.Amount
// 		}
// 	}
// 	return summ
// }
