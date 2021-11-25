package main

import (
	"os"

	controller "uninc/controllers"
	middleware "uninc/middleware"
	routes "uninc/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	// Create Endpoint
	router.POST("/create/coupon", controller.CreateCoupon)
	// List Endpoint
	router.GET("/coupons", controller.ListCoupons)

	router.Run(":" + port)
}
