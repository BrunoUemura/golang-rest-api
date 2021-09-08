package routes

import (
	"github.com/BrunoUemura/golang-rest-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func OrdersRoutes(router *gin.Engine) {
	orders := router.Group("api/v1/orders")
	{
		orders.GET("/", controllers.ShowOrders)
		orders.GET("/:id", controllers.ShowOrder)
		orders.POST("/", controllers.CreateOrder)
		orders.PUT("/", controllers.UpdateOrder)
		orders.DELETE("/:id", controllers.DeleteOrder)
	}
}