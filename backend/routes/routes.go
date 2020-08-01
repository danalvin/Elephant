package routes

import (
	"elephant/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router - returns gin router engine
func Router() *gin.Engine {

	// If we're in production mode, set Gin to "release" mode
	if config.GetConfig().GetString("app.environment") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/")

	// Sub Routes
	v1.GET("/foo", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"foo": "bar"})

		c.Abort()
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Sorry! That one is not handled Here"})
	})

	return router
}
