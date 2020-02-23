package dynamo

import (
	"github.com/btmadison/go/vehicle/pkg/crud"
)

type Repository struct {
	name string
}

func NewRepository() *Repository {
	store := new(Repository)
	store.name = "DynaFoo"
	return store
}

// GetAll gets all
func (m *Repository) GetAllVehicles() []crud.Vehicle {
	vehicles := []crud.Vehicle{}
	return vehicles
}

// GetOneByID returns vehicle with given VIN number
func (m *Repository) GetOneByID(vin string) (crud.Vehicle, error) {
	return crud.Vehicle{
		Vin:        "5e4ef1ad40e8ab28f6e75138",
		Make:       "Acura",
		Model:      "Farger",
		Year:       1969,
		Dealership: "Rent-A-Wreck",
	}, nil
}

// Upsert will Insert or Update existing Vehicle based on globally unique VIN#
func (m *Repository) Upsert(v crud.Vehicle) {
}

// Delete vehicle from in memory inventory
func (m *Repository) Delete(vin string) {
}
