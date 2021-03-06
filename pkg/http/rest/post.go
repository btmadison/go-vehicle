package rest

import (
	"encoding/json"
	"net/http"

	"github.com/btmadison/go-vehicle/pkg/crud"
)

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	body := r.Body
	var v crud.Vehicle

	err := json.NewDecoder(body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	svc.Create(v)
	w.WriteHeader(200)
}
