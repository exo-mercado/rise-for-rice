package consumer

import "github.com/exo-mercado/rise-for-rice/models"

type ConsumerService struct {
	repo ConsumerRepository
}

func NewConsumerService(repo ConsumerRepository) ConsumerService {
	return ConsumerService{
		repo: repo,
	}
}

func (s ConsumerService) Create(consumer models.ConsumerPayload) (models.Consumer, error) {
	return s.repo.Create(consumer)
}