package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	vars := mux.Vars(r)
	vin := vars["vin"]

	vehicle, err := svc.ReadOneByID(vin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	vehicleJSON, _ := json.Marshal(vehicle)

	w.Write(vehicleJSON)
}
