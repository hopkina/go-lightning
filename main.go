package main

import (
	controllers "go-lightning/controllers"
	"go-lightning/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/strikes", controllers.FindStrikes)

	// create
	r.POST("/strikes", controllers.CreateStrike)

	// find by id
	r.GET("/strikes/:id", controllers.FindStrike)

	// update by id
	r.PATCH("/strikes/:id", controllers.UpdateStrike)

	// delete by id
	r.DELETE("/strikes/:id", controllers.DeleteStrike)

	r.Run()
}
