package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	medapp "github.com/mnogohoddovochka/med-app"
)

type successfulLogInResponse struct {
	Token string `json:"token" example:"eyJhbGcjOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHaiOjE2NzY5MjgyMjYsImlhdCI6MTY3Njg4NTAyNiwiZG9jdG9yX2lkIjoxfQ.meE8ccXLe6kJ3gFNLONIZ_PPGOXknQuYac40flbaQ6g"`
}

// @Summary SignUp
// @Tags auth
// @Description Create doctor account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body medapp.Doctor true "Doctor main info"
// @Success 200 {object} newPersonResponse
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input medapp.Doctor

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorisation.CreateDoctor(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type signInInput struct {
	Login    string `json:"login" binding:"required" example:"DrHouse"`
	Password string `json:"password" binding:"required" example:"ilovemedicine777"`
}

// @Summary SignUp
// @Tags auth
// @Description Doctors log in
// @ID login
// @Accept json
// @Produce json
// @Param input body signInInput true "Login details"
// @Success 200 {object} successfulLogInResponse "Returning JWT token"
// @Failure 400,404 {object} errMessage
// @Failure 500 {object} errMessage
// @Failure default {object} errMessage
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorisation.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, successfulLogInResponse{
		Token: token,
	})
}
