package main

import (
	"github.com/btmadison/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/btmadison/go-vehicle/pkg/data/dynamo"
	"github.com/btmadison/btmadison/go-vehicle/pkg/data/inmem"
	"github.com/btmadison/btmadison/go-vehicle/pkg/http/rest"
)

const (
	INMEM = iota
	DYNAMO
)

func main() {
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
