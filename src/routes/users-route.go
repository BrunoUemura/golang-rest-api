package routes

import (
	"github.com/BrunoUemura/golang-rest-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("api/v1/users")
	{
		users.GET("/", controllers.ShowUsers)
		users.GET("/:id", controllers.ShowUser)
		users.POST("/", controllers.CreateUser)
		users.PUT("/", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
		
	}
}
