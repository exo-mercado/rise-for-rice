package client

import "github.com/exo-mercado/rise-for-rice/infrastructure"

type ClientRoute struct {
	Handler infrastructure.GinRouter
	Controller ClientController
}

func NewClientRoute(handler infrastructure.GinRouter, controller ClientController) ClientRoute {
	return ClientRoute{
		Handler: handler,
		Controller: controller,
	}
}

func (r ClientRoute) Setup() {
	client := r.Handler.Gin.Group("/client")
	{
		client.GET("", r.Controller.FindAll)
	}
}