package reservation

import "github.com/exo-mercado/rise-for-rice/models"

type ReservationService struct {
	repo ReservationRepository
}

func NewReservationService(repo ReservationRepository) ReservationService {
	return ReservationService{
		repo: repo,
	}
}

func (s ReservationService) Create(reservation models.ReservationPayload) (models.Reservation, error) {
	return s.repo.Create(reservation)
}

func (s ReservationService) FindByID(id uint) (models.Reservation, error) {
	return s.repo.FindByID(id)
}

func (s ReservationService) FindByGridNumber(gridNumber string) (models.Reservation, error) {
	return s.repo.FindByGridNumber(gridNumber)
}

func (s ReservationService) UsingReserveSpace(id uint) (error) {
	return s.repo.UsingReserveSpace(id)
}

func (s ReservationService) CancelReservation(id uint) (error) {
	return s.repo.CancelReservation(id)
}

func (s ReservationService) FindAll(id string, status string) ([]models.Reservation, error) {
	return s.repo.FindAll(id, status)
}