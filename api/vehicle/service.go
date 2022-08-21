package vehicle

import "github.com/exo-mercado/rise-for-rice/models"

type VehicleService struct {
	repo VehicleRepository
}

func NewVehicleService(repo VehicleRepository) VehicleService {
	return VehicleService{
		repo: repo,
	}
}

func (s VehicleService) Create(vehicle models.VehiclePayload) (models.Vehicle, error) {
	return s.repo.Create(vehicle)
}