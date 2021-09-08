package routes

import (
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	UserRoutes(router)
	BooksRoutes(router)
	OrdersRoutes(router)
	return router
}