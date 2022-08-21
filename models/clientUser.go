package models

import (
	"gorm.io/gorm"
)

type ClientUser struct {
	gorm.Model
	ClientUserFirstName		string	`form:"first_name" json:"first_name"`
	ClientUserLastName		string	`form:"last_name" json:"last_name"`
	ClientUserPhoneNumber	string	`form:"phone_number" json:"phone_number"`
	ClientUserEmail			string	`form:"email" json:"email"`
	ClientUserPosition		string	`form:"position" json:"position"`
	ClientUserPassword		string	`form:"password" json:"password"`
	ClientID				uint	`form:"client_id" json:"client_id"`
	Client					Client
}

func (clientUser *ClientUser) TableName() string {
	return "client_user"
}

//PAYLOADS

type ClientUserPayload struct {
	ClientUserFirstName		string	`form:"first_name" json:"first_name"`
	ClientUserLastName		string	`form:"last_name" json:"last_name"`
	ClientUserPhoneNumber	string	`form:"phone_number" json:"phone_number"`
	ClientUserEmail			string	`form:"email" json:"email"`
	ClientUserPosition		string	`form:"position" json:"position"`
	ClientUserPassword		string	`form:"password" json:"password"`
	ClientID				uint	`form:"client_id" json:"client_id"`
}

//RESPONSES
func (clientUser *ClientUser) BasicResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= clientUser.ID
	resp["first_name"] 			= clientUser.ClientUserFirstName
	resp["last_name"] 			= clientUser.ClientUserLastName
	resp["phone_number"] 		= clientUser.ClientUserPhoneNumber
	resp["email"] 				= clientUser.ClientUserEmail
	resp["position"] 			= clientUser.ClientUserPosition
	resp["created_at"]			= clientUser.CreatedAt
	resp["updated_at"]			= clientUser.UpdatedAt

	return resp
}

func (clientUser *ClientUser) ClientRelationalResponse() map[string]interface{} {
	resp := make(map[string]interface{})

	resp["id"] 					= clientUser.ID
	resp["first_name"] 			= clientUser.ClientUserFirstName
	resp["last_name"] 			= clientUser.ClientUserLastName
	resp["phone_number"] 		= clientUser.ClientUserPhoneNumber
	resp["email"] 				= clientUser.ClientUserEmail
	resp["position"] 			= clientUser.ClientUserPosition
	resp["client"]				= clientUser.Client
	resp["created_at"]			= clientUser.CreatedAt
	resp["updated_at"]			= clientUser.UpdatedAt

	return resp
}
