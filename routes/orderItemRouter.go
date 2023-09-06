package routes

import (
	controller "golang-restaurant/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order_items", controller.GetOrderItems())
	incomingRoutes.GET("/order_items/:order_items_id", controller.GetOrderItem())
	incomingRoutes.GET("/order_items_order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/order_items", controller.CreateOrderItems())
	incomingRoutes.PATCH("/order_items/:order_items_id", controller.UpdateOrderItems())
}
