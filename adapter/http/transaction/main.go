package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/igorgalindop/dio-financial-planning-with-go/model/transaction"
)

func HandleTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch method := r.Method; method {
	case "GET":
		getAllTransactions(w, r)
		break
	case "POST":
		createTransaction(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var layout = "2006-01-02T15:04:05"
	salaryReceived, _ := time.Parse(layout, "2023-11-19T20:00:35")

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1250.0,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	json.NewEncoder(w).Encode(transactions)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var transaction transaction.Transaction
	_ = json.NewDecoder(r.Body).Decode(&transaction)

	fmt.Print(transaction)
}
