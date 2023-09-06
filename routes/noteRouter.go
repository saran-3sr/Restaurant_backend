package routes

import (
	controller "golang-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/note", controller.GetNotes())
	incomingRoutes.GET("/note/:note_id", controller.GetNote())
	incomingRoutes.POST("/note", controller.CreateNote())
	incomingRoutes.PATCH("/note/:note_id", controller.UpdateNote())
}
