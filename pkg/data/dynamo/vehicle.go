package dynamo

import "github.com/btmadison/btmadison/go-vehicle/pkg/crud"

// Vehicle Dyn Resp
type Vehicle struct {
	Pk         string `json:"pk"`
	Sk         string `json:"sk"`
	Make       string `json:"make"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
	Milage     int    `json:"milage"`
	Dealership string `json:"dealership"`
}

// ToCrudVehicle converts a dynamo Vehicle response to a CRUD Vehicle
func (v Vehicle) ToCrudVehicle() crud.Vehicle {
	return crud.Vehicle{
		Vin:        v.Pk,
		Make:       v.Make,
		Model:      v.Model,
		Year:       v.Year,
		Dealership: v.Dealership,
	}
}
