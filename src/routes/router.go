package routes

import (
	"github.com/BrunoUemura/golang-rest-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.ShowUsers)
			users.GET("/:id", controllers.ShowUser)
			users.POST("/", controllers.CreateUser)
			users.PUT("/", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
		orders := main.Group("orders")
		{
			orders.GET("/", controllers.ShowOrders)
			orders.GET("/:id", controllers.ShowOrder)
			orders.POST("/", controllers.CreateOrder)
			orders.PUT("/", controllers.UpdateOrder)
			orders.DELETE("/:id", controllers.DeleteOrder)
		}
	}

	return router
}