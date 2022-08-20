package consumer

import (
	"log"

	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/models"
	"github.com/exo-mercado/rise-for-rice/utils"
)

type ConsumerRepository struct {
	db infrastructure.Database
}

func NewConsumerRepository(db infrastructure.Database) ConsumerRepository {
	return ConsumerRepository{
		db: db,
	}
}

func (r ConsumerRepository) FindAll(consumer models.Consumer, keyword string) (*[]models.Consumer, int64, error) {

	var consumers []models.Consumer
	var totalRows int64 = 0

	queryBuilder := r.db.DB.
		Order("created_at desc").
		Model(&models.Consumer{})

	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("consumer.consumer_first_name LIKE ? ", queryKeyword)).Where("consumer.consumer_last_name LIKE ? ", queryKeyword)
	}

	err := queryBuilder.
		Where(consumer).
		Find(&consumers).
		Count(&totalRows).Error
		
	return &consumers, totalRows, err
}

func (r ConsumerRepository) Create(consumer models.ConsumerPayload) (models.Consumer, error) {

	var consumerModel models.Consumer

	encodedPassword, err := utils.HashPassword(consumer.ConsumerPassword)
	log.Println("Hashing password", encodedPassword)
	if err != nil {
		return consumerModel, err
	}

	consumerModel.ConsumerFirstName 	= consumer.ConsumerFirstName
	consumerModel.ConsumerLastName 		= consumer.ConsumerLastName
	consumerModel.ConsumerPhoneNumber 	= consumer.ConsumerPhoneNumber
	consumerModel.ConsumerPassword 		= encodedPassword

	if err := r.db.DB.Create(&consumerModel).Error; err != nil{
		return consumerModel, err
	}

	return consumerModel, nil

		
}