package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/btmadison/go-vehicle/pkg/crud"
)

var svc crud.Service

// ServeRoutes serves the following routes as a single localhost api
func ServeRoutes(vehicleSvc crud.Service) {
	svc = vehicleSvc
	router := mux.NewRouter()
	registerRoutes(router)
	router.Use(mux.CORSMethodMiddleware(router))
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", router)
}

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/vehicles", list).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/vehicles/{vin}", get).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/vehicles", post).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/vehicles/{vin}", put).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/vehicles/{vin}", delete).Methods(http.MethodDelete, http.MethodOptions)
}
