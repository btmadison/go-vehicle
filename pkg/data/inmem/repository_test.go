package inmem_test

import (
	"testing"

	"github.com/btmadison/go-vehicle/pkg/data/inmem"
)

func TestNewRepository_CreatesSeededRepo(t *testing.T) {
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
