package expense

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
	"time"
)

type storageData struct {
	LastID   int             `json:"last_id"`
	Expenses map[int]Expense `json:"expenses"`
}

type Storage struct {
	mu       sync.Mutex
	lastID   int
	expenses map[int]Expense
	filePath string
}

func NewStorage(filePath string) (*Storage, error) {
	s := &Storage{
		expenses: make(map[int]Expense, 0),
		lastID:   0,
		filePath: filePath,
	}
	if err := s.load(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil
		}
		return nil, err
	}
	return s, nil
}

func (s *Storage) load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	f, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		return nil
	}

	var data storageData
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return err
	}
	s.expenses = data.Expenses
	s.lastID = data.LastID

	return nil
}
func (s *Storage) save() error {

	var data storageData

	data.Expenses = s.expenses
	data.LastID = s.lastID

	tmpFile := s.filePath + ".tmp"

	f, err := os.Create(tmpFile)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "	")

	if err := enc.Encode(&data); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Rename(tmpFile, s.filePath); err != nil {
		return err
	}
	return nil
}

// 1 добавить расход
func (s *Storage) AddExpense(description string, dateString string, amount float64) (Expense, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	exp, err := NewExpense(s.lastID, description, dateString, amount)
	if err != nil {
		s.lastID--
		return Expense{}, err
	}
	s.expenses[s.lastID] = exp

	if err := s.save(); err != nil {
		return Expense{}, err
	}

	return exp, nil
}

// 2 обновить расход
func (s *Storage) UpdateExpense(id int, description string, dateString string, amount float64) (Expense, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.expenses[id]
	if !ok {
		return Expense{}, ErrExpenseNotFound
	}

	exp, err := NewExpense(id, description, dateString, amount)
	if err != nil {
		return Expense{}, err
	}
	s.expenses[id] = exp

	if err := s.save(); err != nil {
		return Expense{}, err
	}

	return exp, nil
}

// 3 удалить расход
func (s *Storage) DeleteExpense(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.expenses[id]
	if !ok {
		return ErrExpenseNotFound
	}

	delete(s.expenses, id)

	if err := s.save(); err != nil {
		return err
	}

	return nil
}

// 4 посмотреть все расходы
func (s *Storage) GetAllExpenses() []Expense {
	s.mu.Lock()
	defer s.mu.Unlock()
	var exps []Expense
	for _, e := range s.expenses {
		exps = append(exps, e)
	}
	return exps
}

// 5 посмотреть сводку раходов
func (s *Storage) TotalSummary() float64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	var summ float64
	for _, e := range s.expenses {
		summ += e.Amount
	}
	return summ

}

// 6 посмотреть сводку расходов по месяцу
func (s *Storage) MonthlySummary(month int) (float64, error) {
	if month > 12 || month < 1 {
		return 0, ErrInvalidMonth
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	var summ float64
	year := time.Now().Year()
	for _, e := range s.expenses {
		if int(e.Date.Month()) == month && year == e.Date.Year() {
			summ += e.Amount
		}
	}
	return summ, nil

}
