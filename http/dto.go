package http

import (
	"encoding/json"
	"expense-tracker/expense"
	"time"
)

//data transfer object

type ExpenseDTO struct {
	Description string
	Date        string
	Amount      float64
}

func (e ExpenseDTO) ExpenseDataValidate() error {
	if e.Amount <= 0 {
		return expense.ErrInvalidAmount
	}
	if e.Description == "" {
		return expense.ErrInvalidDescription
	}
	if e.Date != "" {
		_, err := time.Parse("2006-01-02", e.Date)
		if err != nil {
			return expense.ErrInvalidData
		}
	}
	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
func NewErrDTO(message string) ErrorDTO {
	return ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
}
