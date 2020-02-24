package dynamo

import "github.com/btmadison/btmadison/go-vehicle/pkg/crud"

// Vehicle Dyn Resp
type Vehicle struct {
	Pk         string
	Sk         string
	Make       string
	Model      string
	Year       int
	Dealership string
}

func (v Vehicle) ToCrudVehicle() crud.Vehicle {
	return crud.Vehicle{
		Vin:        v.Pk,
		Make:       v.Make,
		Model:      v.Model,
		Year:       v.Year,
		Dealership: v.Dealership,
	}
}
