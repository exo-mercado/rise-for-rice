package client

import (
	"net/http"

	"github.com/exo-mercado/rise-for-rice/models"
	"github.com/exo-mercado/rise-for-rice/utils"
	"github.com/gin-gonic/gin"
)

type ClientController struct {
	service ClientService
}

func NewClientController(service ClientService) ClientController {
	return ClientController{
		service: service,
	}
}

func (c ClientController) FindAll(g *gin.Context) {
	var clients models.Client

	keyword := g.Query("keyword")

	data, total, err := c.service.FindAll(clients, keyword)

	if err != nil {
		utils.ErrorJSON(g, http.StatusInternalServerError, err)
	}

	responseArray := make([]map[string]interface{}, 0)

	for _, client := range *data {
		resp := client.AreasRelationalResponse()
		responseArray = append(responseArray, resp)
	}

	g.JSON(http.StatusOK, gin.H{
		"content": responseArray,
		"totalElements":   total,
	})
}
