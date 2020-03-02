package main

import (
	"log"

	"github.com/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/go-vehicle/pkg/data/dynamo"
	"github.com/btmadison/go-vehicle/pkg/data/inmem"
	"github.com/btmadison/go-vehicle/pkg/http/rest"
	"github.com/joho/godotenv"
)

const (
	INMEM = iota
	DYNAMO
)

func main() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var svc crud.Service

	switch db := DYNAMO; db {
	case DYNAMO:
		repo := dynamo.NewRepository()
		svc = crud.NewService(repo)
	case INMEM:
		repo := inmem.NewRepository()
		svc = crud.NewService(repo)
	default:
		panic("INVALID DATA SOURCE")
	}

	rest.ServeRoutes(svc)
}
