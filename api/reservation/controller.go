package reservation

import (
	"net/http"

	"github.com/exo-mercado/rise-for-rice/models"
	"github.com/exo-mercado/rise-for-rice/utils"
	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	service ReservationService
}

func NewReservationController(service ReservationService) ReservationController {
	return ReservationController{
		service: service,
	}
}

func (c ReservationController) Create(g *gin.Context) {
	
	var reservation models.ReservationPayload
	if err := g.ShouldBindJSON(&reservation); err != nil {
		utils.ErrorJSON(g, http.StatusBadRequest, err)
		return
	}

	response, err := c.service.Create(reservation)
	if err != nil {
		utils.ErrorJSON(g, http.StatusInternalServerError, err)
		return
	}

	g.JSON(http.StatusOK, response.BasicResponse())
}

func (c ReservationController) FindAll(g *gin.Context) {
	id := g.Query("id")
	status := g.Query("status")
	response, err := c.service.FindAll(id, status)
	if err != nil {
		utils.ErrorJSON(g, http.StatusInternalServerError, err)
		return
	}

	responseArray := make([]map[string]interface{}, 0)

	for _, reservationResponse := range response {
		resp := reservationResponse.BasicResponse()
		responseArray = append(responseArray, resp)
	}

	g.JSON(http.StatusOK, gin.H{
		"content": responseArray,
		"totalElements":   len(responseArray),
	})
}