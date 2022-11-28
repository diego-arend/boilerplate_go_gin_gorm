package routes

import (
	"github.com/diego-arend/boilerplate_go_gin_gorm/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
		}
		cars := main.Group("cars")
		{
			cars.GET("/", controllers.ListCars)
			cars.GET("/:id", controllers.ShowCar)
			cars.POST("/", controllers.CreateCar)
			cars.PUT("/", controllers.UpdateCar)
			cars.DELETE("/:id", controllers.DeleteCar)
		}
	}

	return router
}
