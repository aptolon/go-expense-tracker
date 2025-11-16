package expense

import "errors"

var (
	ErrInvalidDescription = errors.New("описание не может быть пустым")
	ErrInvalidData        = errors.New("формат даты YYYY-MM-DD")
	ErrInvalidAmount      = errors.New("сумма не может быть отрицательной")
	ErrExpenseNotFound    = errors.New("такая покупка не найдена")
)
