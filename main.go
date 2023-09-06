package main

import (
	"golang-restaurant/database"
	"golang-restaurant/middleware"
	"golang-restaurant/routes"
	"os"

	"github.com/gin-gonic/gin"
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
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TablesRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.InvoiceRoutes(router)
	routes.NoteRoutes(router)
	routes.UserRoutes(router)
}
