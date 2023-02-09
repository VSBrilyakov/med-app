package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	medapp "github.com/mnogohoddovochka/med-app"
)

type getAllDoctorsResponse struct {
	Data []medapp.Doctor `json:"data"`
}

func (h *Handler) getAllDoctors(c *gin.Context) {
	doctors, err := h.services.DoctorList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDoctorsResponse{
		Data: doctors,
	})
}

func (h *Handler) getDoctorById(c *gin.Context) {

}
