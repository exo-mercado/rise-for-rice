package models

import (
	"gorm.io/gorm"
)

type Consumer struct {
	gorm.Model
	ConsumerFirstName		string	`form:"first_name" json:"first_name"`
	ConsumerLastName		string	`form:"last_name" json:"last_name"`
	ConsumerPhoneNumber		string	`form:"phone_numer" json:"phone_number"`
	ConsumerPassword		string	`form:"password" json:"password"`
	Vehicles				[]Vehicle
}

func (consumer *Consumer) TableName() string {
	return "consumer"
}

//Payload
type ConsumerPayload struct {
	ConsumerFirstName		string	`gorm:"required" form:"first_name" json:"first_name"`
	ConsumerLastName		string	`gorm:"required" form:"last_name" json:"last_name"`
	ConsumerPhoneNumber		string	`gorm:"required" form:"phone_number" json:"phone_number"`
	ConsumerPassword		string	`gorm:"required" form:"password" json:"password"`
}


//RESPONSES
func (consumer *Consumer) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= consumer.ID
	resp["first_name"] 			= consumer.ConsumerFirstName
	resp["last_name"] 			= consumer.ConsumerLastName
	resp["phone_number"] 		= consumer.ConsumerPhoneNumber
	resp["created_at"]			= consumer.CreatedAt
	resp["updated_at"]			= consumer.UpdatedAt

	return resp
}

func (consumer *Consumer) VehicleRelationalResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= consumer.ID
	resp["first_name"] 			= consumer.ConsumerFirstName
	resp["last_name"] 			= consumer.ConsumerLastName
	resp["phone_number"] 		= consumer.ConsumerPhoneNumber
	resp["vehicles"]			= consumer.Vehicles
	resp["created_at"]			= consumer.CreatedAt
	resp["updated_at"]			= consumer.UpdatedAt

	return resp
}