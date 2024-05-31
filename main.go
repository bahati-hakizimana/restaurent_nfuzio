package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"restaurant_management/database"
	"restaurant_management/routes"
	"restaurant_management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {

	port := os.Getenv("PORT")
     
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authantication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.TableRoutes(router)
	routes.InvoiceRoutes(router)

	router.run(":" +port)
}