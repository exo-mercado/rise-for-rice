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

func (s ConsumerService) Login(consumer models.LoginPayload) (models.Consumer, error) {

	
	return s.repo.Login(consumer)
}

func (s ConsumerService) FindByID(id int, join string) (models.Consumer, error) {
	return s.repo.FindByID(id, join)
}