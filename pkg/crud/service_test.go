package crud_test

import (
	"errors"
	"testing"

	"github.com/btmadison/go-vehicle/pkg/crud"
)

type mockRepo struct {
}

var mockRepository mockRepo

func TestUnit_CanCreateNewService(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	if svc == nil {
		t.Error("service creation failed when passed valid repository")
	}
}

func TestUnit_SvcReadAll_CallsRepoGetAll(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	result, err := svc.ReadAll()
	if err != nil {
		t.Error("Error calling repo get all")
	}
	if result[0].Vin != "MOCKVIN1" {
		t.Error("Error calling through to repo get all func")
	}
}

func TestUnit_SvcGet_CallsRepoGet(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	result, err := svc.ReadOneByID("test_foo_id")
	if err != nil {
		t.Error("Error calling repo get 1")
	}
	if result.Vin != "test_foo_id" {
		t.Error("Error calling through to repo get 1 func")
	}
}

func TestUnit_SvcGet_HandlesError(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	_, err := svc.ReadOneByID("THROW_ERR")
	if err.Error() != "MOCK_GET1_ERROR" {
		t.Error("Failed to return an error when it should have")
	}
}

func TestUnit_SvcDelete_CallsRepoDelete(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	err := svc.Delete("FOO_TEST")
	if err != nil {
		t.Error("Error calling repo delete")
	}
}

func TestUnit_SvcDelete_CallsRepoDeleteRepoError(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	err := svc.Delete("THROW_ERR")
	if err == nil {
		t.Error("Should have returned an error")
	}
}

func TestUnit_ServiceCreate_CallsRepoUpsert(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	v, err := svc.Create(crud.Vehicle{Vin: "FOO"})
	if err != nil {
		t.Error("create vehicle err")
	}
	if v.Vin != "FOO" {
		t.Error("Failed to pass vehicle object to repo and handle repsonse")
	}
}

func TestUnit_ServiceUpdate_CallsRepoUpsert_WhenRouteVinMatchesPayloadVin(t *testing.T) {
	mockRepository = mockRepo{}
	svc := crud.NewService(mockRepository)
	v, err := svc.Update("foo_id", crud.Vehicle{Vin: "foo_id"})
	if err != nil {
		t.Error("error, should have called through to repo upsert")
	}
	if v.Vin != "foo_id" {
		t.Error("Error on service update vehicle call to repo")
	}
}

func (m mockRepo) GetAllVehicles() ([]crud.Vehicle, error) {
	return []crud.Vehicle{{Vin: "MOCKVIN1"}}, nil
}

func (m mockRepo) GetOneByID(vin string) (crud.Vehicle, error) {
	if vin == "THROW_ERR" {
		return crud.Vehicle{}, errors.New("MOCK_GET1_ERROR")
	}
	return crud.Vehicle{Vin: vin}, nil
}

func (m mockRepo) Delete(vin string) error {
	if vin == "THROW_ERR" {
		err := errors.New("MOCK_DELETE_ERROR")
		return err
	}
	return nil
}

func (m mockRepo) Upsert(vehicle crud.Vehicle) error {
	if vehicle.Vin == "THROW_ERR" {
		return errors.New("MOCK_UPSERT_ERROR")
	}
	return nil
}
