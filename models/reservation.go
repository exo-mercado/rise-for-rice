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
	Status 			string	`form:"status" json:"status"`
	Vehicle			Vehicle
}

func (reservation *Reservation) TableName() string {
	return "reservation"
}


//RESPONSES
func (reservation *Reservation) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= reservation.ID
	resp["from"] 				= reservation.ReservationFrom
	resp["to"] 					= reservation.ReservationTo
	resp["grid_number"] 		= reservation.GridNumber
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
	resp["vehicle"]				= reservation.Vehicle
	resp["created_at"]			= reservation.CreatedAt
	resp["updated_at"]			= reservation.UpdatedAt

	return resp
}

