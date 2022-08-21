package reservation

import (
	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/models"
)


type ReservationRepository struct {
	db infrastructure.Database
}

func NewReservationRepository(db infrastructure.Database) ReservationRepository {
	return ReservationRepository{
		db: db,
	}
}

func (r ReservationRepository) Create(reservation models.ReservationPayload) (models.Reservation, error) {
	var reservationModel models.Reservation

	reservationModel.ReservationFrom 	= reservation.ReservationFrom
	reservationModel.ReservationTo 		= reservation.ReservationTo
	reservationModel.VehicleID 			= reservation.VehicleID
	reservationModel.GridNumber			= reservation.GridNumber
	reservationModel.AreaID				= reservation.AreaID
	reservationModel.Status				= "RESERVED"

	err := r.db.DB.Create(&reservationModel).Error
	if err != nil {
		return reservationModel, err
	}

	return reservationModel, nil
}

func (r ReservationRepository) FindByID(id uint) (models.Reservation, error) {
	var reservation models.Reservation

	err := r.db.DB.Find(&reservation, "id = ?", id).Error
	if err != nil {
		return reservation, err
	}

	return reservation, nil
}

func (r ReservationRepository) FindByGridNumber(gridNumber string) (models.Reservation, error) {
	var reservation models.Reservation

	err := r.db.DB.Find(&reservation, "grid_number = ?", gridNumber).Error
	if err != nil {
		return reservation, err
	}

	return reservation, nil
}

func (r ReservationRepository) UsingReserveSpace(id uint) (error) {
	err := r.db.DB.Find(&models.Reservation{}).Where("id = ?", id).Update("status", "ONSITE").Error
	if err != nil {
		return err
	}

	return nil
}

func (r ReservationRepository) FindByUserID(id uint) (error) {
	err := r.db.DB.Find(&models.Reservation{}).Where("consumer_id = ?", id).Update("status", "ONSITE").Error
	if err != nil {
		return err
	}

	return nil
}

func (r ReservationRepository) CancelReservation(id uint) (error) {
	var reservation models.Reservation

	r.db.DB.Find(&reservation, "id = ?", id)


	err := r.db.DB.Model(&models.Reservation{}).Where("id = ?", id).Update("status", "CANCELLED").Error
	if err != nil {
		return err
	}

	return nil
}

func (r ReservationRepository) FindAll(id string, status string) ([]models.Reservation, error) {
	var reservations []models.Reservation

	queryBuilder := r.db.DB.
		Preload("Vehicle.Consumer").
		Order("created_at desc").
		Model(&models.Reservation{})

	if status != "" {
		queryBuilder = queryBuilder.Where("status = ?", status)
	}

	if id != "" {
		queryBuilder = queryBuilder.Where("vehicle_id IN (?)",r.db.DB.Table("vehicle").Select("ID").Where("consumer_id = ?", id))
	}

	err := queryBuilder.Find(&reservations).Error

	return reservations, err
}