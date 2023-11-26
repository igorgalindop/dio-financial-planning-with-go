package http

import (
	"net/http"

	"github.com/igorgalindop/dio-financial-planning-with-go/adapter/http/actuator"
	"github.com/igorgalindop/dio-financial-planning-with-go/adapter/http/transaction"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Init comment
func Init() {
	http.HandleFunc("/", transaction.HandleTransactions)

	http.HandleFunc("/health", actuator.Health)

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
