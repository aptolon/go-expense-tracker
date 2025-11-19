package expense

import "errors"

var (
	ErrInvalidDescription = errors.New("description cannot be empty")
	ErrInvalidData        = errors.New("date format must be YYYY-MM-DD")
	ErrInvalidAmount      = errors.New("amount cannot be negative")
	ErrExpenseNotFound    = errors.New("expense not found")
	ErrInvalidMonth       = errors.New("invalid month")
)
