package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody comment
type HealthBody struct {
	Status string `json:"status"`
}

// Health comment
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)

}
