package rest

import (
	"encoding/json"
	"net/http"

	"github.com/btmadison/go-vehicle/pkg/crud"
	"github.com/gorilla/mux"
)

func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	body := r.Body
	var v crud.Vehicle

	vars := mux.Vars(r)
	vin := vars["vin"]

	decodeErr := json.NewDecoder(body).Decode(&v)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	_, err := svc.Update(vin, v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(200)
}
