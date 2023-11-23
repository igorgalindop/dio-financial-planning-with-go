package http

import (
	"net/http"

	"github.com/igorgalindop/dio-financial-planning-with-go/adapter/http/actuator"
	"github.com/igorgalindop/dio-financial-planning-with-go/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/", transaction.HandleTransactions)

	http.HandleFunc("/health", actuator.Health)

	http.ListenAndServe(":8080", nil)
}
