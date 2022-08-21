package area

import (
	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/models"
)

type AreaRepository struct {
	db infrastructure.Database
}

func NewAreaRepository(db infrastructure.Database) AreaRepository {
	return AreaRepository{
		db: db,
	}
}

func (r AreaRepository) Create(area models.AreaPayload) (models.Area, error) {
	var areaModel models.Area
	areaModel.AreaName 		= area.AreaName
	areaModel.NumberOfGrid 	= area.NumberOfGrid
	areaModel.AreaCapacity	= area.AreaCapacity
	areaModel.ClientID 		= area.ClientID	
	areaModel.IsActive 		= true

	err := r.db.DB.Create(&areaModel).Error
	if err != nil {
		return areaModel, err
	}

	return areaModel, nil
}

