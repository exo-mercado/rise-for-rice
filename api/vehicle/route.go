package vehicle

import "github.com/exo-mercado/rise-for-rice/infrastructure"

type VehicleRoute struct {
	Handler infrastructure.GinRouter
	Controller VehicleController
}

func NewVehicleRoute(handler infrastructure.GinRouter, controller VehicleController) VehicleRoute {
	return VehicleRoute{
		Handler: handler,
		Controller: controller,
	}
}

func (r VehicleRoute) Setup() {
	vehicle := r.Handler.Gin.Group("/vehicle")
	{
		vehicle.POST("", r.Controller.Create)
	}
}