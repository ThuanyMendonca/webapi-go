package routes

import (
	"github.com/ThuanyMendonca/webapi-go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		login := main.Group("/login")
		{
			login.POST("/", controllers.Login)
		}
		user := main.Group("/user")
		{
			user.POST("/", controllers.CreateUser)
		}
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowAllBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}
