package vehicle

import (
	"net/http"
	"strconv"

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

func (c VehicleController) FindByConsumerID(g *gin.Context) {
	consumerID, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid consumer id"})
		return
	}

	response, err := c.service.FindByConsumerID(consumerID)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	if len(response) == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "No vehicle found"})
		return
	}

	responseArray := make([]map[string]interface{}, 0)

	for _, vehicle := range response {
		resp := vehicle.BasicResponse()
		responseArray = append(responseArray, resp)
	}

	g.JSON(http.StatusOK, gin.H{
		"content": responseArray,
		"totalElements":   len(response),
	})
}