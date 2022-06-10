package main

import (
	"os"

	routes "github.com/Kaustubh8691/golang-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted"})
	})
	router.GET("/api2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted in 2"})
	})

	router.Run(":" + port)
}
