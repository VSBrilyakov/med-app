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

// @Summary Create visit
// @Security ApiKeyAuth
// @Tags visits
// @Description Add a patient visit into database
// @ID create-visit
// @Accept  json
// @Produce  json
// @Param input body medapp.Visit true "Visit main info"
// @Success 200 {object} newPersonResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/visits [post]
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

	c.JSON(http.StatusOK, newPersonResponse{
		Id: id,
	})
}

// @Summary Get All visits
// @Security ApiKeyAuth
// @Tags visits
// @Description Get all visits information
// @ID get-all-visits
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllVisitsResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/visits [get]
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

// @Summary Get visit By Id
// @Security ApiKeyAuth
// @Tags visits
// @Description Get visit information by id
// @ID get-visit-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Visit ID"
// @Success 200 {object} medapp.VisitOutput
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/visits/:id [get]
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

// @Summary Update visit
// @Security ApiKeyAuth
// @Tags visits
// @Description update visit information
// @ID update-visit
// @Accept  json
// @Produce  json
// @Param id path int true "Visit ID"
// @Param input body medapp.UpdVisit true "New visit info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/visits/:id [put]
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

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

// @Summary Delete visit
// @Security ApiKeyAuth
// @Tags visits
// @Description Delete visit information from database
// @ID delete-visit
// @Accept  json
// @Produce  json
// @Param id path int true "Visit ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/visits/:id [delete]
func (h *Handler) deleteVisit(c *gin.Context) {
	_, err := getDoctorId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.VisitList.DeleteVisit(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
