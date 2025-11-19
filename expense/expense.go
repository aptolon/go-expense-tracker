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
