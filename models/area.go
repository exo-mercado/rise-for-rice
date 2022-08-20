package models

import (
	"gorm.io/gorm"
)

type Area struct {
	gorm.Model
	AreaName		string  `form:"name" json:"name"`
	AreaCapacity	uint    `form:"capacity" json:"capacity"`
	IsActive		bool	`form:"is_active" json:"is_active"`
	NumberOfGrid	uint	`form:"number_of_grid" json:"number_of_grid"`
	ClientID		uint	`form:"client_id" json:"client_id"`
}

func (area *Area) TableName() string {
	return "area"
}

//RESPONSES
func (area *Area) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= area.ID
	resp["name"] 				= area.AreaName
	resp["capacity"] 			= area.AreaCapacity
	resp["is_active"] 			= area.IsActive
	resp["number_of_grid"] 		= area.NumberOfGrid
	resp["created_at"]			= area.CreatedAt
	resp["updated_at"]			= area.UpdatedAt

	return resp
}

