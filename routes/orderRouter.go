package routes


import (
	"restaurant_management/controller"

	"github.com/gin-gonic/gin"
)



func OrderRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/orders", controller.GetOrders())
	incomingRoutes.GET("/orderss/:order_id", controller.GetOrder())
	incomingRoutes.POST("/order", controller.CreateOrder())
	incomingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}