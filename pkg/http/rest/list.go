package rest

import (
	"encoding/json"
	"net/http"
)

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	vehicles := svc.ReadAll()
	vehiclesJSON, _ := json.Marshal(vehicles)
	w.Write(vehiclesJSON)
}
