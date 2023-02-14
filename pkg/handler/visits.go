package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	medapp "github.com/mnogohoddovochka/med-app"
)

type getAllVisitsResponse struct {
	Data []medapp.VisitOutput `json:"data"`
}

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

func (h *Handler) getAllVisits(c *gin.Context) {
	visits, err := h.services.VisitList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllVisitsResponse{
		Data: visits,
	})
}

func (h *Handler) getVisitById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	visit, err := h.services.VisitList.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, visit)
}

func (h *Handler) updateVisit(c *gin.Context) {
	_, err := getDoctorId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input medapp.UpdVisit
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateVisit(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}
