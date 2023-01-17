package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	medapp "github.com/mnogohoddovochka/med-app"
)

func (h *Handler) signUp(c *gin.Context) {
	var input medapp.Doctor

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorisation.CreateDoctor(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
