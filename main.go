package main

import (
	controllers "go-lightning/controllers"
	_ "go-lightning/docs"
	"go-lightning/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Lightning Strikes API
// @version         1.0
// @description     Lightning strikes API

// @host            localhost:8080
// @BasePath        /api/v1

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1 := r.Group("/api/v1")
	{
		strikes := v1.Group("/strikes")
		{
			strikes.GET(":id", controllers.FindStrike)
			strikes.GET("", controllers.FindStrikes)
			strikes.POST("", controllers.CreateStrike)
			strikes.PATCH(":id", controllers.UpdateStrike)
			strikes.DELETE(":id", controllers.DeleteStrike)

		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
