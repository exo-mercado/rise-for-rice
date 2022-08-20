package consumer

import (
	"net/http"

	"github.com/exo-mercado/rise-for-rice/models"
	"github.com/exo-mercado/rise-for-rice/utils"
	"github.com/gin-gonic/gin"
)

type ConsumerController struct {
	service ConsumerService
}

func NewConsumerController(service ConsumerService) ConsumerController {
	return ConsumerController{
		service: service,
	}
}

func (c ConsumerController) Create(g *gin.Context) {

	var consumer models.ConsumerPayload
	if err := g.BindJSON(&consumer); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		utils.ErrorJSON(g, http.StatusBadRequest, err)
		return
	}

	response, err := c.service.Create(consumer)
	if err != nil {

		utils.ErrorJSON(g, http.StatusInternalServerError, err)
		return
	}

	g.JSON(http.StatusOK, response)
}
