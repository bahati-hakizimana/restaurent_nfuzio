package routes

import (
	"restaurant_management/controller"

	"github.com/gin-gonic/gin"
)

func TableRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tabless/:table_id", controller.GetTable())
	incomingRoutes.POST("/table", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
}
