package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/henriquesml/go-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.POST("/", controllers.CreateBook)
			books.GET("/", controllers.ShowBooks)
			books.GET("/:id", controllers.ShowBook)
		}
	}

	return router
}
