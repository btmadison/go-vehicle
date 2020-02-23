package inmem

import (
	"errors"

	"github.com/btmadison/btmadison/go-vehicle/pkg/crud"
)

type Repository struct {
	vehicleMap map[string]crud.Vehicle
}

// NewRepository returns a repository pointer with seeded fake in memory vehicle inventory
func NewRepository() *Repository {
	store := new(Repository)
	populateSeedData(store)
	return store
}

// GetAll gets all vehicles from the inmem repo
func (m *Repository) GetAllVehicles() []crud.Vehicle {
	vehicles := []crud.Vehicle{}
	for _, value := range m.vehicleMap {
		vehicles = append(vehicles, value)
	}

	return vehicles
}

// GetOneByID returns vehicle with given VIN number from the inmem repo
func (m *Repository) GetOneByID(vin string) (crud.Vehicle, error) {
	v, exists := m.vehicleMap[vin]
	if exists == false {
		err := errors.New("vehicle not found")
		return v, err
	}

	return v, nil
}

// Upsert will Insert or Update existing Vehicle in the inmem repo map
func (m *Repository) Upsert(v crud.Vehicle) {
	m.vehicleMap[v.Vin] = v
}

// Delete vehicle from inmem repo map
func (m *Repository) Delete(vin string) {
	delete(m.vehicleMap, vin)
}

func populateSeedData(store *Repository) {
	store.vehicleMap = make(map[string]crud.Vehicle)

	store.vehicleMap["5e4ef1ad40e8ab28f6e75138"] = crud.Vehicle{
		Vin:        "5e4ef1ad40e8ab28f6e75138",
		Make:       "Acura",
		Model:      "Farger",
		Year:       1969,
		Dealership: "Rent-A-Wreck",
	}

	store.vehicleMap["5e4ef1adb58e48ac9a1aa176"] = crud.Vehicle{
		Vin:        "5e4ef1adb58e48ac9a1aa176",
		Make:       "Ford",
		Model:      "Mustang",
		Year:       1982,
		Dealership: "ShadySales",
	}

	store.vehicleMap["5e4ef1adf5de4ddafa2ed9d5"] = crud.Vehicle{
		Vin:        "5e4ef1adf5de4ddafa2ed9d5",
		Make:       "Ford",
		Model:      "Model T",
		Year:       1921,
		Dealership: "Scamalot",
	}

	store.vehicleMap["5e4ef1ad066871df6bae98a9"] = crud.Vehicle{
		Vin:        "5e4ef1ad066871df6bae98a9",
		Make:       "Chevy",
		Model:      "Impalla",
		Year:       1996,
		Dealership: "ShadySales",
	}

	store.vehicleMap["5e4ef1ad08165005d2aa1742"] = crud.Vehicle{
		Vin:        "5e4ef1ad08165005d2aa1742",
		Make:       "Audi",
		Model:      "A8",
		Year:       2014,
		Dealership: "BMad Motors",
	}
}
