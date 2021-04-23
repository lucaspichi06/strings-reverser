package app

import "github.com/gin-gonic/gin"

// mapRoutes to create the routes
func (c *controllers) mapRoutes(router *gin.Engine) {
	router.GET("/ping", c.status.HandlePing)
	router.POST("/revert_string", c.revert.HandleReversion)
}
