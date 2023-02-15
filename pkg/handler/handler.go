package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mnogohoddovochka/med-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.doctorIdentity)
	{
		doctors := api.Group("/doctors")
		{
			doctors.GET("/", h.getAllDoctors)
			doctors.GET("/:id", h.getDoctorById)
		}

		patients := api.Group("/patients")
		{
			patients.POST("/", h.createPatient)
			patients.GET("/", h.getAllPatients)
			patients.GET("/:id", h.getPatientById)
			patients.PUT("/:id", h.updatePatient)
			patients.DELETE("/:id", h.deletePatient)
		}

		visits := api.Group("/visits")
		{
			visits.POST("/", h.createVisit)
			visits.GET("/", h.getAllVisits)
			visits.GET("/:id", h.getVisitById)
			visits.PUT("/:id", h.updateVisit)
			visits.DELETE("/:id", h.deleteVisit)
		}
	}

	return router
}
