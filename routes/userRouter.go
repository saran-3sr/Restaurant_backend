package routes
import (
	"github.com/gin-gonic/gin"
	"golang-restaurant/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controllers.GetUser())
	incomingRoutes.GET("/users/:user_id",controllers.GetUsers())
	incomingRoutes.POST("/users/signup",controllers.Signup())
	incomingRoutes.POST("/users/login",controllers.Login())

}