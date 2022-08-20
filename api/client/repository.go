package client

import (
	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/models"
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
		Preload("Areas").
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