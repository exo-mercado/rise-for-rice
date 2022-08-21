package vehicle

import (
	"net/http"

	"github.com/exo-mercado/rise-for-rice/models"
	"github.com/exo-mercado/rise-for-rice/utils"
	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	service VehicleService
}

func NewVehicleController(service VehicleService) VehicleController {
	return VehicleController{
		service: service,
	}
}

func (c VehicleController) Create(g *gin.Context) {
	
	var vehicle models.VehiclePayload
	if err := g.ShouldBindJSON(&vehicle); err != nil {
		utils.ErrorJSON(g, http.StatusBadRequest, err)
		return
	}

	response, err := c.service.Create(vehicle)
	if err != nil {
		utils.ErrorJSON(g, http.StatusInternalServerError, err)
		return
	}

	g.JSON(http.StatusOK, response.BasicResponse())
}