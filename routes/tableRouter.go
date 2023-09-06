package routes

import (
	controller "golang-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func TablesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:tables_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTables())
	incomingRoutes.PATCH("/tables/:tables_id", controller.UpdateTables())
}
