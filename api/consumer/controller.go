package consumer

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/exo-mercado/rise-for-rice/models"
	"github.com/exo-mercado/rise-for-rice/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	consumer := models.ConsumerPayload{}

	if err := g.ShouldBindJSON(&consumer); err != nil {
		utils.ErrorJSON(g, http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	fmt.Println("consumer", consumer)

	response, err := c.service.Create(consumer)
	if err != nil {
		utils.ErrorJSON(g, http.StatusInternalServerError, err)
		return
	}

	g.JSON(http.StatusOK, response.BasicResponse() )
}

func (c ConsumerController) Login(g *gin.Context) {

	consumer := models.LoginPayload{}
	secret := []byte(os.Getenv("JWT_SECRET"))

	if err := g.ShouldBindJSON(&consumer); err != nil {
		utils.ErrorJSON(g, http.StatusBadRequest, err)
		return
	}

	response, err := c.service.Login(consumer)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": response.ID,	
		"phone_number": response.ConsumerPhoneNumber,
		
	})

	signed, err := token.SignedString(secret)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"token": signed})

}


func (c ConsumerController) FindByID(g *gin.Context) {

	join := g.Query("join")

	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		utils.ErrorJSON(g, http.StatusInternalServerError, err)
		return
	}

	response, err := c.service.FindByID(id, join)
	if err != nil {
		fmt.Println(err)
		g.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Cannot fetch to server"})
		return
	}

	if response.ID == 0 {
		g.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Cannot find the consumer"})
		return
	}

	g.JSON(http.StatusOK, response.VehicleRelationalResponse() )
}