package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	medapp "github.com/mnogohoddovochka/med-app"
)

type getAllPatientsResponse struct {
	Data []medapp.Patient `json:"data"`
}

func (h *Handler) createPatient(c *gin.Context) {
	_, err := getDoctorId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input medapp.Patient
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreatePatient(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getAllPatients(c *gin.Context) {
	patients, err := h.services.PatientList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllPatientsResponse{
		Data: patients,
	})
}

func (h *Handler) getPatientById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	patient, err := h.services.PatientList.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, patient)
}
