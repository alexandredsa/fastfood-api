package main

import (
	"fastfood-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	setupRoutes()
}

func setupRoutes() {
	router := gin.Default()

	api := router.Group("/api")
	reviewsRoute := api.Group("/types/:id/reviews")
	foodReviewController := controllers.FoodReviewController{Router: reviewsRoute}
	foodReviewController.SetupRoutes()
	typesRoute := api.Group("/types/")
	foodSummaryController := controllers.FoodSummaryController{Router: typesRoute}
	foodSummaryController.SetupRoutes()
	router.Run(":8080")
}
