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
	VehiclePlate	string	`form:"plate" json:"plate"`
	IsDefault		bool	`form:"is_default" json:"is_default"`
	ConsumerID		uint	`form:"consumer_id" json:"consumer_id"`
	Consumer		Consumer
}

func (vehicle *Vehicle) TableName() string {
	return "vehicle"
}

//Payload
type VehiclePayload struct {
	VehicleMake		string	`form:"make" json:"make" binding:"required"`
	VehicleModel	string	`form:"model" json:"model" binding:"required"`
	VehicleYear		string	`form:"year" json:"year" binding:"required"`
	VehicleColor	string	`form:"color" json:"color" binding:"required"`
	VehiclePlate	string	`form:"plate" json:"plate" binding:"required"`
	ConsumerID		uint	`form:"consumer_id" json:"consumer_id" binding:"required"`
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