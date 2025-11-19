package main

import (
	"expense-tracker/expense"
	"expense-tracker/http"
	"fmt"
)

func main() {
	storage, err := expense.NewStorage("./str.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	httpHandlers := http.NewHTTPHandlers(storage)
	httpServer := http.NewHTTPServer(httpHandlers)
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
