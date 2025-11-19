package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.
		Path("/expenses").
		Methods("POST").
		HandlerFunc(s.httpHandlers.HandlerAddExpense)
	router.
		Path("/expenses/{id}").
		Methods("PATCH").
		HandlerFunc(s.httpHandlers.HandlerUpdateExpense)
	router.
		Path("/expenses/{id}").
		Methods("DELETE").
		HandlerFunc(s.httpHandlers.HandlerDeleteExpense)
	router.
		Path("/expenses").
		Methods("GET").
		HandlerFunc(s.httpHandlers.HandlerGetAllExpenses)
	router.
		Path("/expenses/summary").
		Methods("GET").
		Queries("month", "{month}").
		HandlerFunc(s.httpHandlers.HandlerMonthlySummary)
	router.
		Path("/expenses/summary").
		Methods("GET").
		HandlerFunc(s.httpHandlers.HandlerTotalSummary)
	fmt.Println("сервер запущен на :8080")
	return http.ListenAndServe(":8080", router)
}
