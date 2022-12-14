package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	ReservationFrom	string	`form:"from" json:"from"`
	ReservationTo	string	`form:"to" json:"to"`
	GridNumber		uint	`form:"grid_number" json:"grid_number"`
	VehicleID		uint	`form:"vehicle_id" json:"vehicle_id"`
	AreaID			uint	`form:"area_id" json:"area_id"`
	Status 			string	`form:"status" json:"status"`
	Vehicle			Vehicle
	Area			Area
}

func (reservation *Reservation) TableName() string {
	return "reservation"
}

//PAYLOAD
type ReservationPayload struct {
	ReservationFrom	string	`form:"from" json:"from" binding:"required"`
	ReservationTo	string	`form:"to" json:"to" binding:"required"`
	GridNumber		uint	`form:"grid_number" json:"grid_number" binding:"required"`
	VehicleID		uint	`form:"vehicle_id" json:"vehicle_id" binding:"required"`
	AreaID  		uint	`form:"area_id" json:"area_id" binding:"required"`
}

//RESPONSES
func (reservation *Reservation) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= reservation.ID
	resp["from"] 				= reservation.ReservationFrom
	resp["to"] 					= reservation.ReservationTo
	resp["grid_number"] 		= reservation.GridNumber
	resp["status"]				= reservation.Status
	resp["created_at"]			= reservation.CreatedAt
	resp["updated_at"]			= reservation.UpdatedAt

	return resp
}

func (reservation *Reservation) VehicleRelationalResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= reservation.ID
	resp["from"] 				= reservation.ReservationFrom
	resp["to"] 					= reservation.ReservationTo
	resp["grid_number"] 		= reservation.GridNumber
	resp["vehicle"]				= reservation.Vehicle.BasicResponse()
	resp["status"]				= reservation.Status
	resp["created_at"]			= reservation.CreatedAt
	resp["updated_at"]			= reservation.UpdatedAt

	return resp
}

