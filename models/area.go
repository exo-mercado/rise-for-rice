package models

import (
	"github.com/exo-mercado/rise-for-rice/utils"
	"gorm.io/gorm"
)

type Area struct {
	gorm.Model
	AreaName		string  `form:"name" json:"name"`
	AreaCapacity	uint    `form:"capacity" json:"capacity"`
	IsActive		bool	`form:"is_active" json:"is_active"`
	NumberOfGrid	uint	`form:"number_of_grid" json:"number_of_grid"`
	ClientID		uint	`form:"client_id" json:"client_id"`
	Reservations	[]Reservation
}

func (area *Area) TableName() string {
	return "area"
}

//PAYLOAD
type AreaPayload struct {
	AreaName		string  `form:"name" json:"name" binding:"required"`
	AreaCapacity	uint    `form:"capacity" json:"capacity" binding:"required"`
	NumberOfGrid	uint	`form:"number_of_grid" json:"number_of_grid" binding:"required"`
	ClientID		uint	`form:"client_id" json:"client_id" binding:"required"`
}

//RESPONSES
func (area *Area) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})
	occupantCount := 0

	for _, reservation := range area.Reservations {
		if utils.UnixTimeIsNotExpired(reservation.ReservationTo){
			occupantCount++
		}
	}

	resp["id"] 					= area.ID
	resp["name"] 				= area.AreaName
	resp["capacity"] 			= area.AreaCapacity
	resp["is_active"] 			= area.IsActive
	resp["number_of_grid"] 		= area.NumberOfGrid
	resp["occupant"]			= occupantCount
	resp["created_at"]			= area.CreatedAt
	resp["updated_at"]			= area.UpdatedAt

	return resp
}

