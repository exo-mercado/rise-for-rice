package consumer

import "github.com/exo-mercado/rise-for-rice/infrastructure"

type ConsumerRoute struct {
	Handler infrastructure.GinRouter
	Controller ConsumerController
}

func NewConsumerRoute(handler infrastructure.GinRouter, controller ConsumerController) ConsumerRoute {
	return ConsumerRoute{
		Handler: handler,
		Controller: controller,
	}
}

func (r ConsumerRoute) Setup() {
	consumer := r.Handler.Gin.Group("/consumer")
	{
		consumer.GET("/:id", r.Controller.FindByID)
		consumer.POST("", r.Controller.Create)
		consumer.POST("/login", r.Controller.Login)
	}
}