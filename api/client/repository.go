package client

import (
	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/models"
	"github.com/exo-mercado/rise-for-rice/utils"
)

type ClientRepository struct {
	db infrastructure.Database
}

func NewClientRepository(db infrastructure.Database) ClientRepository {
	return ClientRepository{
		db: db,
	}
}

func (r ClientRepository) FindAll(client models.Client, keyword string) (*[]models.Client, int64, error) {

	var clients []models.Client
	var totalRows int64 = 0

	queryBuilder := r.db.DB.
		Preload("Areas.Reservations").
		Order("created_at desc").
		Model(&models.Client{})

	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("client.client_name LIKE ? ", queryKeyword))
	}

	err := queryBuilder.
		Where(client).
		Find(&clients).
		Count(&totalRows).Error
		
	return &clients, totalRows, err
}

func (r ClientRepository) CreateClientUser(client models.ClientUserPayload) (models.ClientUser, error) {

	var clientUser models.ClientUser

	hashPassword, _ := utils.HashPassword(client.ClientUserPassword)

	clientUser.ClientUserEmail			= client.ClientUserEmail
	clientUser.ClientUserFirstName 		= client.ClientUserFirstName
	clientUser.ClientUserLastName		= client.ClientUserLastName
	clientUser.ClientUserPhoneNumber 	= client.ClientUserPhoneNumber
	clientUser.ClientUserPosition 		= client.ClientUserPosition
	clientUser.ClientUserPassword 		= hashPassword
	
	err := r.db.DB.Create(&clientUser).Error
	if err != nil {
		return clientUser, err
	}

	return clientUser, nil
}