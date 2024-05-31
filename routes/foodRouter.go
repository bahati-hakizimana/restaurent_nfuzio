package routes

import (
	"restaurant_management/controller"

	"github.com/gin-gonic/gin"
)


func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Get("/foods", controller.GetFoods())
	incomingRoutes.Get("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}