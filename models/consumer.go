package models

import (
	"gorm.io/gorm"
)

type Consumer struct {
	gorm.Model
	ConsumerFirstName		string	`json:"first_name"`
	ConsumerLastName		string	`json:"last_name"`
	ConsumerPhoneNumber		string	`gorm:"unique" json:"phone_number"`
	ConsumerPassword		string	`json:"password"`
	Vehicles				[]Vehicle
}

func (consumer *Consumer) TableName() string {
	return "consumer"
}

//Payload
type ConsumerPayload struct {
	ConsumerFirstName		string	`json:"first_name" binding:"required"`
	ConsumerLastName		string	`json:"last_name" binding:"required"`
	ConsumerPhoneNumber		string	`json:"phone_number" binding:"required"`
	ConsumerPassword		string	`json:"password" binding:"required"`
}


type LoginPayload struct {
	ConsumerPhoneNumber		string	`json:"phone_number" binding:"required"`
	ConsumerPassword		string	`json:"password" binding:"required"`
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