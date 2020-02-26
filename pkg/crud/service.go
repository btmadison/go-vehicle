package crud

import "errors"

// Repository is injected into the service to provide data access that corresponds to the service crud functionality
type Repository interface {
	GetAllVehicles() ([]Vehicle, error)
	GetOneByID(vin string) (Vehicle, error)
	Delete(vin string)
	Upsert(vehicle Vehicle)
}

// Service defines the methods for vehicle crud access
type Service interface {
	Create(vehicle Vehicle) (Vehicle, error)
	ReadAll() ([]Vehicle, error)
	ReadOneByID(vin string) (Vehicle, error)
	Update(vin string, vehicle Vehicle) (Vehicle, error)
	Delete(vin string) error
}

type service struct {
	repo Repository
}

// NewService creats a new service with injected repository
func NewService(r Repository) Service {
	return &service{r}
}

func (svc service) ReadAll() ([]Vehicle, error) {
	return svc.repo.GetAllVehicles()
}

func (svc *service) ReadOneByID(vin string) (Vehicle, error) {
	return svc.repo.GetOneByID(vin)
}

func (svc service) Delete(vin string) error {
	svc.repo.Delete(vin)
	return nil
}

func (svc service) Create(vehicle Vehicle) (Vehicle, error) {
	svc.repo.Upsert(vehicle)
	return vehicle, nil
}

func (svc service) Update(vin string, vehicle Vehicle) (Vehicle, error) {
	if vin != vehicle.Vin {
		err := errors.New("forbidden operation")
		return vehicle, err
	}
	svc.repo.Upsert(vehicle)
	return vehicle, nil
}
