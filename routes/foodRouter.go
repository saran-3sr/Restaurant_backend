package routes

import (
	controller "golang-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFood())
	incomingRoutes.GET("/foods/:food_id", controller.GetFoods())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id", controller.UpdateFoods())
}
