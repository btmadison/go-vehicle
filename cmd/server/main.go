package main

import (
	"log"
	"net/http"

	"github.com/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/go-vehicle/pkg/data/dynamo"
	"github.com/btmadison/go-vehicle/pkg/data/inmem"
	"github.com/btmadison/go-vehicle/pkg/http/rest"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	INMEM = iota
	DYNAMO
)

func main() {
	godotenv.Load("local.env")
	var svc crud.Service
	var repo crud.Repository

	switch db := DYNAMO; db {
	case DYNAMO:
		repo = dynamo.NewRepository()
	case INMEM:
		repo = inmem.NewRepository()
	default:
		panic("INVALID DATA SOURCE")
	}

	svc = crud.NewService(repo)

	router := mux.NewRouter()
	rest.RegisterRoutes(router, svc)
	log.Fatal(http.ListenAndServe(":8080", router))
}
