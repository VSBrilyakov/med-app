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

// @Summary Create patient
// @Security ApiKeyAuth
// @Tags patients
// @Description Add a patient information into database
// @ID create-patient
// @Accept  json
// @Produce  json
// @Param input body medapp.Patient true "Patient main info"
// @Success 200 {object} newPersonResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/patients [post]
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

	c.JSON(http.StatusOK, newPersonResponse{
		Id: id,
	})
}

// @Summary Get All patients
// @Security ApiKeyAuth
// @Tags patients
// @Description Get all patients information
// @ID get-all-patients
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllPatientsResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/patients [get]
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

// @Summary Get patient By Id
// @Security ApiKeyAuth
// @Tags patients
// @Description Get patient information by id
// @ID get-patient-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} medapp.Patient
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/patients/:id [get]
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

// @Summary Update patient
// @Security ApiKeyAuth
// @Tags patients
// @Description update patient info
// @ID update-patient
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Param input body medapp.UpdPatient true "New patient info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/patients/:id [put]
func (h *Handler) updatePatient(c *gin.Context) {
	_, err := getDoctorId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input medapp.UpdPatient
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdatePatient(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete patient
// @Security ApiKeyAuth
// @Tags patients
// @Description Delete patient information from database
// @ID delete-patient
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /api/patients/:id [delete]
func (h *Handler) deletePatient(c *gin.Context) {
	_, err := getDoctorId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.PatientList.DeletePatient(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
