package setup

import (
	"github.com/exo-mercado/rise-for-rice/api/client"
	"github.com/exo-mercado/rise-for-rice/api/consumer"
	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/models"
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



	// migrating User model to datbase table
	db.DB.AutoMigrate(
		&models.Area{},
		&models.Client{},
		&models.ClientUser{},
		&models.Consumer{},
		&models.Reservation{},
		&models.Vehicle{},
	);
}
