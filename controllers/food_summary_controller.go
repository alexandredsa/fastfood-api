package controllers

import (
	"fastfood-api/models"
	"fastfood-api/repositories"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type FoodSummaryController struct {
	Router                *gin.RouterGroup
	foodSummaryRepository repositories.FoodSummaryRepository
}

func (fsr *FoodSummaryController) SetupRoutes() {
	fsr.foodSummaryRepository = repositories.FoodSummaryRepository{}
	fsr.list()
	fsr.save()
}

func (fsr *FoodSummaryController) list() {
	fsr.Router.GET("/", func(c *gin.Context) {
		foodSummaries := fsr.foodSummaryRepository.FindAll()
		c.JSON(200, foodSummaries)
	})
}

func (fsr *FoodSummaryController) save() {
	fsr.Router.POST("/", func(c *gin.Context) {
		var foodSummary models.FoodSummary
		if err := c.ShouldBindJSON(&foodSummary); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		foodSummary.ID = primitive.NewObjectID()
		fsr.foodSummaryRepository.Save(foodSummary)
	})
}
