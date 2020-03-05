package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/btmadison/go-vehicle/pkg/crud"
)

var svc crud.Service

// RegisterRoutes with the handler, injecting router and crud service
func RegisterRoutes(r *mux.Router, vehicleSvc crud.Service) {
	svc = vehicleSvc
	r.HandleFunc("/vehicles", list).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/vehicles/{vin}", get).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/vehicles", post).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/vehicles/{vin}", put).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/vehicles/{vin}", delete).Methods(http.MethodDelete, http.MethodOptions)
}
