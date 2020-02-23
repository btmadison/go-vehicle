package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Delete handler deletes a single item by id
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)
	vin := vars["vin"]
	err := svc.Delete(vin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(200)
}
