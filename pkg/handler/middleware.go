package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorisationHeader = "Authorization"
	doctorCtxId         = "doctorId"
)

func (h *Handler) doctorIdentity(c *gin.Context) {
	header := c.GetHeader(authorisationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[1] == "" {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	doctorId, err := h.services.Authorisation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(doctorCtxId, doctorId)
}

func getDoctorId(c *gin.Context) (int, error) {
	id, ok := c.Get(doctorCtxId)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "doctor id is not found")
		return 0, errors.New("doctor id is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "doctor id is of invalid type")
		return 0, errors.New("doctor id is of invalid type")
	}

	return idInt, nil
}
