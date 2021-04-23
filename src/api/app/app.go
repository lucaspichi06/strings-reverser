package app

import (
	"github.com/gin-gonic/gin"
)

// Start the application
func Start() *gin.Engine {
	controllers := load()
	return start(controllers)
}

func start(controllers *controllers) *gin.Engine {
	router := gin.Default()

	controllers.mapRoutes(router)

	return router
}
