package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	medapp "github.com/mnogohoddovochka/med-app"
)

func (h *Handler) createVisit(c *gin.Context) {
	var input medapp.Visit
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateVisit(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
