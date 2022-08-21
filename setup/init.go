package setup

import (
	"github.com/exo-mercado/rise-for-rice/api/client"
	"github.com/exo-mercado/rise-for-rice/api/consumer"
	"github.com/exo-mercado/rise-for-rice/api/reservation"
	"github.com/exo-mercado/rise-for-rice/api/vehicle"
	"github.com/exo-mercado/rise-for-rice/infrastructure"
)

func InitializeServices(router infrastructure.GinRouter) {
	db := infrastructure.NewDatabase() // databse has been initialized and configured

	// client
	clientRepository :=  client.NewClientRepository(db)
	clientService := client.NewClientService(clientRepository)  
	clientController := client.NewClientController(clientService)
	clientRoute := client.NewClientRoute(router, clientController)
	clientRoute.Setup()

	// consumer
	consumerRepository :=  consumer.NewConsumerRepository(db)
	consumerService := consumer.NewConsumerService(consumerRepository)  
	consumerController := consumer.NewConsumerController(consumerService)
	consumerRoute := consumer.NewConsumerRoute(router, consumerController)
	consumerRoute.Setup()

	// vehicle
	vehicleRepository :=  vehicle.NewVehicleRepository(db)
	vehicleService := vehicle.NewVehicleService(vehicleRepository)
	vehicleController := vehicle.NewVehicleController(vehicleService)
	vehicleRoute := vehicle.NewVehicleRoute(router, vehicleController)
	vehicleRoute.Setup()

	//reservation

	reservationRepository := reservation.NewReservationRepository(db)
	reservationService := reservation.NewReservationService(reservationRepository)
	reservationController := reservation.NewReservationController(reservationService)
	reservationRoute := reservation.NewReservationRoute(router, reservationController)
	reservationRoute.Setup()



	// migrating User model to datbase table
	// db.DB.AutoMigrate(
	// 	&models.Area{},
	// 	&models.Client{},
	// 	&models.ClientUser{},
	// 	&models.Consumer{},
	// 	&models.Reservation{},
	// 	&models.Vehicle{},
	// );
}
