package orders

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// NewHandler receives its database before it creates any clients.
func NewHandler(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", func(w http.ResponseWriter, r *http.Request) {
		var order struct {
			ID int `json:"id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&order); err != nil || order.ID <= 0 {
			http.Error(w, "invalid order", http.StatusBadRequest)
			return
		}
		if _, err := db.ExecContext(r.Context(), "INSERT INTO orders(id) VALUES ($1)", order.ID); err != nil {
			http.Error(w, "cannot create order", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
	return mux
}
