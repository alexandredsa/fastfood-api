package controllers

import (
	"fastfood-api/models"
	"fastfood-api/repositories"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type FoodReviewController struct {
	Router                *gin.RouterGroup
	foodReviewRepository  repositories.FoodReviewRepository
	foodSummaryRepository repositories.FoodSummaryRepository
}

func (frc *FoodReviewController) SetupRoutes() {
	frc.foodReviewRepository = repositories.FoodReviewRepository{}
	frc.foodSummaryRepository = repositories.FoodSummaryRepository{}
	frc.list()
	frc.save()
}

func (frc *FoodReviewController) list() {
	frc.Router.GET("/", func(c *gin.Context) {
		foodSummaryId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		foodReviews := frc.foodReviewRepository.FindAllByTypeId(foodSummaryId)
		c.JSON(200, foodReviews)
	})
}

func (frc *FoodReviewController) save() {
	frc.Router.POST("/", func(c *gin.Context) {
		foodSummaryId, _ := primitive.ObjectIDFromHex(c.Param("id"))
		var foodReview models.FoodReview
		foodSummary := models.FoodSummary{ID: foodSummaryId}
		if err := c.ShouldBindJSON(&foodReview); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		foodReview.FoodSummaryId = foodSummaryId
		if err := frc.foodSummaryRepository.UpdateSummary(&foodSummary); err != nil {
			c.Status(400)
			return
		}
		frc.foodReviewRepository.Save(&foodReview)
		c.Status(201)
	})
}
