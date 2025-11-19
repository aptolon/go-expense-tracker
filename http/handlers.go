package http

import (
	"encoding/json"
	"errors"
	"expense-tracker/expense"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	Storage *expense.Storage
}

func NewHTTPHandlers(Storage *expense.Storage) *HTTPHandlers {
	return &HTTPHandlers{
		Storage: Storage,
	}
}

// 1 добавить расход
/*
pattern: 	/expenses
method:		POST
info: 		JSON in HTTP request body

success:
	- status code: 		201 Created
	- response body: 	JSON represent created expense
failed:
	- status code: 		400,409,500...
	- response body: 	JSON with error + time
*/
func (h *HTTPHandlers) HandlerAddExpense(w http.ResponseWriter, r *http.Request) {
	var expDTO ExpenseDTO
	if err := json.NewDecoder(r.Body).Decode(&expDTO); err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := expDTO.ExpenseDataValidate(); err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	exp, err := h.Storage.AddExpense(expDTO.Description, expDTO.Date, expDTO.Amount)
	if err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		return
	}
	b, err := json.MarshalIndent(exp, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

// 2 обновить расход
/*
pattern: 	/expenses/{id}
method:		PATCH
info: 		pattern + JSON in HTTP request body

success:
	- status code: 		200 OK
	- response body: 	JSON represent changed expense
failed:
	- status code: 		400,404,409,500...
	- response body: 	JSON with error + time
*/
func (h *HTTPHandlers) HandlerUpdateExpense(w http.ResponseWriter, r *http.Request) {
	var expDTO ExpenseDTO
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&expDTO); err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := expDTO.ExpenseDataValidate(); err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	exp, err := h.Storage.UpdateExpense(id, expDTO.Description, expDTO.Date, expDTO.Amount)
	if err != nil {
		errDTO := NewErrDTO(err.Error())
		if errors.Is(err, expense.ErrExpenseNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(exp, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusContinue)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

// 3 удалить расход
/*
pattern: 	/expenses/{id}
method:		DELETE
info: 		pattern

success:
	- status code: 		204 No Content
	- response body: 	-
failed:
	- status code: 		400,404,500...
	- response body: 	JSON with error + time
*/
func (h *HTTPHandlers) HandlerDeleteExpense(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	if err := h.Storage.DeleteExpense(id); err != nil {
		errDTO := NewErrDTO(err.Error())
		if errors.Is(err, expense.ErrExpenseNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// 4 посмотреть все расходы
/*
pattern: 	/expenses
method:		GET
info: 		-

success:
	- status code: 		200 OK
	- response body: 	JSON represented found expenses
failed:
	- status code: 		400,500...
	- response body: 	JSON with error + time
*/
func (h *HTTPHandlers) HandlerGetAllExpenses(w http.ResponseWriter, r *http.Request) {
	exp := h.Storage.GetAllExpenses()
	b, err := json.MarshalIndent(exp, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusContinue)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

// 5 посмотреть сводку раходов
/*
pattern: 	/expenses/summary
method:		GET
info: 		-

success:
	- status code: 		200 OK
	- response body: 	JSON represent total summary
failed:
	- status code: 		400,500...
	- response body: 	JSON with error + time
*/
func (h *HTTPHandlers) HandlerTotalSummary(w http.ResponseWriter, r *http.Request) {
	exp := h.Storage.TotalSummary()
	b, err := json.MarshalIndent(exp, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusContinue)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

// 6 посмотреть сводку расходов по месяцу
/*
pattern: 	/expenses/summary?month={month}
method:		GET
info: 		query params

success:
	- status code: 		200 OK
	- response body: 	JSON represent changed expense
failed:
	- status code: 		400,404,500...
	- response body: 	JSON with error + time
*/
func (h *HTTPHandlers) HandlerMonthlySummary(w http.ResponseWriter, r *http.Request) {
	month, err := strconv.Atoi(mux.Vars(r)["month"])
	if err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	exp, err := h.Storage.MonthlySummary(month)
	if err != nil {
		errDTO := NewErrDTO(err.Error())
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
	b, err := json.MarshalIndent(exp, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusContinue)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}
