package routes

import (
	"github.com/BrunoUemura/golang-rest-api/src/controllers"
	"github.com/gin-gonic/gin"
)

func BooksRoutes(router *gin.Engine) {
	books := router.Group("api/v1/books")
	{
		books.GET("/", controllers.ShowBooks)
		books.GET("/:id", controllers.ShowBook)
		books.POST("/", controllers.CreateBook)
		books.PUT("/", controllers.UpdateBook)
		books.DELETE("/:id", controllers.DeleteBook)
	}
}