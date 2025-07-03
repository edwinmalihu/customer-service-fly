package route

import (
	"customer-service/controller"
	"customer-service/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// SetupRoutes : all the routes are defined here
func SetupRoutes() {
	httpRouter := gin.Default()
	httpRouter.Use(middleware.CORSMiddleware())

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	customerController := controller.NewCustomerController()

	apiRoutes := httpRouter.Group("/api")

	{
		apiRoutes.POST("/login", customerController.LoginUser)
	}

	httpRouter.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
	// httpRouter.Run(":8082")
}
