package router

import (
	"github.com/Kayky18/GK_API/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHanler()

	// Initialize routes here
	v1 := router.Group("/api/v1")
	{
		//Create Patients
		v1.POST("/patient/create", handler.CreatePatients)

		//Show a Patient
		v1.GET("/patient/:id", handler.ShowPatients)

		//List Patients
		v1.GET("/patients", handler.ListPatients)

		//Delete Patients
		v1.DELETE("/patient/:id", handler.DeletePatients)

		//Update Patients
		v1.PUT("/patient/:id", handler.UpdatePatients)

		v1.Any("/patient", handler.PatientBadRequest)

	}

}
