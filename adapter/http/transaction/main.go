package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/igorgalindop/dio-financial-planning-with-go/model/transaction"
	"github.com/igorgalindop/dio-financial-planning-with-go/util"
)

// HandleTransactions comment
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

	var transactions = transaction.Transactions{
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1250.0,
			Type:      0,
			CreatedAt: util.StringToTime("2023-11-10T00:00:00"),
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
