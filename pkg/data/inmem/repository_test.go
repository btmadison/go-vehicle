package inmem_test

import (
	"testing"

	"github.com/btmadison/go-vehicle/pkg/data/inmem"
)

func TestUnit_NewRepository_CreatesSeededRepo(t *testing.T) {
	r := inmem.NewRepository()
	vehicles, err := r.GetAllVehicles()
	if err != nil {
		t.Error(err)
	}
	count := len(vehicles)
	if count != 5 {
		t.Errorf("Should Seed with 5 fake vehicles, got %d instead", count)
	}
}

func TestUnit_DeleteItem_WhenExists_ErrWhenDoesntExist(t *testing.T) {
	r := inmem.NewRepository()
	_, err := r.GetOneByID("5e4ef1ad40e8ab28f6e75138")
	if err != nil {
		t.Error(err)
	}
	err = r.Delete("5e4ef1ad40e8ab28f6e75138")
	if err != nil {
		t.Error(err.Error())
	}
	v, err := r.GetOneByID("5e4ef1ad40e8ab28f6e75138")
	if err == nil || v.Vin != "" {
		t.Error("Deleted item should no longer exist")
	}
}

func TestUnit_GetOne_WhenExists(t *testing.T) {
	r := inmem.NewRepository()
	v, err := r.GetOneByID("5e4ef1ad40e8ab28f6e75138")
	if err != nil {
		t.Error(err)
	}
	if v.Vin != "5e4ef1ad40e8ab28f6e75138" {
		t.Errorf("did not get existing vehicle with correct vin")
	}
}

func TestUnit_GetOne_WhenDoesntExists(t *testing.T) {
	r := inmem.NewRepository()
	_, err := r.GetOneByID("ABCDEFG")
	if err.Error() != "vehicle not found" {
		t.Fail()
	}
}
