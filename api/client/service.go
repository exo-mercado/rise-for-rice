package client

import "github.com/exo-mercado/rise-for-rice/models"

type ClientService struct {
	repo ClientRepository
}

func NewClientService(repo ClientRepository) ClientService {
	return ClientService{
		repo: repo,
	}
}

func (s ClientService) FindAll(client models.Client, keyword string) (*[]models.Client, int64, error) {
	return s.repo.FindAll(client, keyword)
}
