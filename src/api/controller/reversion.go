package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucaspichi06/strings-reverter/src/api/domain"
	"github.com/lucaspichi06/strings-reverter/src/api/errors"
	"log"
	"net/http"
)

type reversionService interface {
	Revert(*domain.ReversionRequest) (*domain.ReversionResponse, error)
}

type reversion struct {
	sReversion reversionService
}

func NewReversionController(sReversion reversionService) *reversion {
	return &reversion{
		sReversion: sReversion,
	}
}

func (r *reversion) HandleReversion(c *gin.Context) {
	body := &domain.ReversionRequest{}
	if err := c.BindJSON(body); err != nil {
		log.Println("[controller.reversion.HandleReversion] content has invalid format", err)
		c.JSON(http.StatusBadRequest, errors.NewBadRequestAppError(fmt.Sprintf("content has invalid format: %s", err)))
		return
	}

	resp, err := r.sReversion.Revert(body)
	if err != nil {
		log.Println("[controller.reversion.HandleReversion] error while reverting the message", err)
		c.JSON(http.StatusInternalServerError, errors.NewInternalServerAppError("error while reverting the message", err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
