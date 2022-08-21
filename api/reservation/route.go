package reservation

import "github.com/exo-mercado/rise-for-rice/infrastructure"

type ReservationRoute struct {
	Handler infrastructure.GinRouter
	Controller ReservationController
}

func NewReservationRoute(handler infrastructure.GinRouter, controller ReservationController) ReservationRoute {
	return ReservationRoute{
		Handler: handler,
		Controller: controller,
	}
}

func (r ReservationRoute) Setup() {
	reservation := r.Handler.Gin.Group("/reservation")
	{
		reservation.POST("", r.Controller.Create)
		reservation.GET("", r.Controller.FindAll)
	}
}