package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaspichi06/strings-reverter/src/api/controller"
	"github.com/lucaspichi06/strings-reverter/src/api/service"
)

type revertController interface {
	HandleReversion(*gin.Context)
}

type statusController interface {
	HandlePing(*gin.Context)
}

type controllers struct {
	revert revertController
	status statusController
}

func load() *controllers {
	// Application layer
	reversionService := service.NewReversionService()

	// UI layer
	reversionController := controller.NewReversionController(reversionService)
	statusController := controller.NewStatusController()

	controllers := &controllers{
		revert: reversionController,
		status: statusController,
	}

	return controllers
}
