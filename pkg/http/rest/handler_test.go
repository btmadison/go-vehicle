package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/btmadison/go-vehicle/pkg/crud"
	"github.com/btmadison/go-vehicle/pkg/http/rest"
	"github.com/gorilla/mux"
)

type mockSvc struct {
}

var mockService mockSvc
var mRouter *mux.Router

func setupTestRoutes() {
	mockService = mockSvc{}
	mRouter = mux.NewRouter()
	rest.RegisterRoutes(mRouter, mockService)
	httptest.NewServer(mRouter)
}

func TestUnit_NonExistantRoute_Returns404(t *testing.T) {
	setupTestRoutes()
	req, _ := http.NewRequest("GET", "http://localhost:8080/BADROUTE", nil)
	resp := httptest.NewRecorder()
	mRouter.ServeHTTP(resp, req)
	if resp.Code != 404 {
		t.Error("non existant route did not return 404 err code")
	}
}

func TestUnit_HandlerRestList_GetsAllVehicles(t *testing.T) {
	setupTestRoutes()
	req, _ := http.NewRequest("GET", "http://localhost:8080/vehicles", nil)
	resp := httptest.NewRecorder()
	mRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Error("GET vehicles route not defined")
	}
	body := resp.Body.String()
	b := []byte(body)

	vehicles := []crud.Vehicle{}
	json.Unmarshal(b, &vehicles)

	if vehicles[0].Vin != "TESTVEHICLEVIN1" {
		t.Error("List Route failed to handle request")
	}
}

func TestUnit_HandlerRestDelete(t *testing.T) {
	setupTestRoutes()
	req, _ := http.NewRequest("DELETE", "http://localhost:8080/vehicles/1234", nil)
	resp := httptest.NewRecorder()
	mRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Error("DELETE vehicles route failed")
	}
}

func TestUnit_HandlerRestPut(t *testing.T) {
	setupTestRoutes()
	req, _ := http.NewRequest("PUT", "http://localhost:8080/vehicles/abc", bytes.NewBuffer([]byte(`{"Vin":"abc"}`)))
	resp := httptest.NewRecorder()
	mRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Error("PUT vehicles route failed")
	}
}

func TestUnit_HandlerRestPost(t *testing.T) {
	setupTestRoutes()
	req, _ := http.NewRequest("POST", "http://localhost:8080/vehicles", bytes.NewBuffer([]byte(`{"Vin":"abc"}`)))
	resp := httptest.NewRecorder()
	mRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Error("POST vehicle route failed")
	}
}

func (svc mockSvc) ReadAll() ([]crud.Vehicle, error) {
	vs := []crud.Vehicle{}
	vs = append(vs, crud.Vehicle{
		Vin:        "TESTVEHICLEVIN1",
		Make:       "Acura",
		Model:      "MDX",
		Year:       1969,
		Dealership: "FOO DEALERSHIP 1",
	})
	return vs, nil
}

func (svc mockSvc) ReadOneByID(vin string) (crud.Vehicle, error) {
	return crud.Vehicle{Vin: vin}, nil
}

func (svc mockSvc) Delete(vin string) error {
	return nil
}

func (svc mockSvc) Create(vehicle crud.Vehicle) (crud.Vehicle, error) {
	return vehicle, nil
}

func (svc mockSvc) Update(vin string, vehicle crud.Vehicle) (crud.Vehicle, error) {
	return crud.Vehicle{Vin: vin, Dealership: vehicle.Dealership}, nil
}
