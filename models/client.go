package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ClientName                   	string    `form:"name" json:"name"`
	ClientLatitude               	float64 `form:"latitude" json:"latitude"`
	ClientLongitude              	float64 `form:"longitude" json:"longitude"`
	ClientReservationFee         	float64 `form:"reservation_fee" json:"reservation_fee"`
	ClientReservationDateAndTime 	string  `form:"reservation_time" json:"reservation_time"`
	ClientInitialHours           	int64	`form:"initial_hours" json:"initial_hours"`
	Areas							[]Area
}

func (client *Client) TableName() string {
	return "client"
}

//RESPONSES
func (client *Client) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= client.ID
	resp["name"] 				= client.ClientName
	resp["latitude"] 			= client.ClientLatitude
	resp["longitude"]			= client.ClientLongitude
	resp["reservation_fee"]		= client.ClientReservationFee
	resp["reservation_time"]	= client.ClientReservationDateAndTime
	resp["initial_hours"] 		= client.ClientInitialHours
	resp["created_at"]			= client.CreatedAt
	resp["updated_at"]			= client.UpdatedAt

	return resp
}

func (client *Client) AreasRelationalResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= client.ID
	resp["name"] 				= client.ClientName
	resp["latitude"] 			= client.ClientLatitude
	resp["longitude"]			= client.ClientLongitude
	resp["reservation_fee"]		= client.ClientReservationFee
	resp["reservation_time"]	= client.ClientReservationDateAndTime
	resp["initial_hours"] 		= client.ClientInitialHours
	resp["areas"]				= client.Areas
	resp["created_at"]			= client.CreatedAt
	resp["updated_at"]			= client.UpdatedAt

	return resp
}
