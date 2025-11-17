package expense

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type storageData struct {
	LastID   int       `json:"last_id"`
	Expenses []Expense `json:"expenses"`
}

type Storage struct {
	mu       sync.Mutex
	lastID   int
	expenses []Expense
	filePath string
}

func NewStorage(filePath string) (*Storage, error) {
	s := &Storage{
		expenses: make([]Expense, 0),
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

	var data storageData

	if err := json.NewDecoder(f).Decode(&data); err != nil {
		return err
	}
	s.expenses = data.Expenses
	s.lastID = data.LastID

	return nil
}
func (s *Storage) save() error {
	s.mu.Lock()
	defer s.mu.Unlock()

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
	s.expenses = append(s.expenses, exp)

	if err := s.save(); err != nil {
		return Expense{}, err
	}

	return exp, nil
}

// 2 обновить расход
// 3 удалить расход
// 4 посмотреть все расходы
// 5 посмотреть сводку раходов
// 6 посмотреть сводку расходов по месяцу
