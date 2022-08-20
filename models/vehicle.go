package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	VehicleMake		string	`form:"make" json:"make"`
	VehicleModel	string	`form:"model" json:"model"`
	VehicleYear		string	`form:"year" json:"year"`
	VehicleColor	string	`form:"color" json:"color"`
	IsDefault		bool	`form:"is_default" json:"is_default"`
	ConsumerID		uint	`form:"consumer_id" json:"consumer_id"`
	Consumer		Consumer
}

func (vehicle *Vehicle) TableName() string {
	return "vehicle"
}


//RESPONSES
func (vehicle *Vehicle) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= vehicle.ID
	resp["make"] 				= vehicle.VehicleMake
	resp["model"] 				= vehicle.VehicleModel
	resp["year"] 				= vehicle.VehicleYear
	resp["color"] 				= vehicle.VehicleColor
	resp["created_at"]			= vehicle.CreatedAt
	resp["updated_at"]			= vehicle.UpdatedAt

	return resp
}

func (vehicle *Vehicle) ConsumerRelationalResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= vehicle.ID
	resp["make"] 				= vehicle.VehicleMake
	resp["model"] 				= vehicle.VehicleModel
	resp["year"] 				= vehicle.VehicleYear
	resp["color"] 				= vehicle.VehicleColor
	resp["consumer"]			= vehicle.Consumer
	resp["created_at"]			= vehicle.CreatedAt
	resp["updated_at"]			= vehicle.UpdatedAt

	return resp
}