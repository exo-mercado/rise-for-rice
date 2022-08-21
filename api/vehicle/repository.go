package vehicle

import (
	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/models"
)


type VehicleRepository struct {
	db infrastructure.Database
}

func NewVehicleRepository(db infrastructure.Database) VehicleRepository {
	return VehicleRepository{
		db: db,
	}
}

func (r VehicleRepository) Create(vehicle models.VehiclePayload) (models.Vehicle, error) {
	var vehicleModel models.Vehicle

	r.db.DB.Model(&models.Vehicle{}).Where("is_default = ?", true).Update("is_default", false)

	vehicleModel.VehicleMake 	= vehicle.VehicleMake
	vehicleModel.VehicleModel 	= vehicle.VehicleModel
	vehicleModel.VehicleYear 	= vehicle.VehicleYear
	vehicleModel.VehicleColor 	= vehicle.VehicleColor
	vehicleModel.VehiclePlate 	= vehicle.VehiclePlate
	vehicleModel.ConsumerID 	= vehicle.ConsumerID
	vehicleModel.IsDefault 		= true

	err := r.db.DB.Create(&vehicleModel).Error
	if err != nil {
		return vehicleModel, err
	}

	return vehicleModel, nil
}