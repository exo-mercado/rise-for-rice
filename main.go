package main

import (
	// "github.com/gin-contrib/cors"
	"github.com/exo-mercado/rise-for-rice/infrastructure"
	"github.com/exo-mercado/rise-for-rice/setup"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter() //router has been initialized and configured
	setup.InitializeServices(router)        //passing router
	router.Gin.Run(":8080")                 //server started on 8000 port

	// router.Gin.Use(cors.Default())
}
