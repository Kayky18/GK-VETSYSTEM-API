package router

import (
	docs "github.com/Kayky18/GK_API/docs"
	"github.com/Kayky18/GK_API/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {
	handler.InitializeHanler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	// Initialize routes here
	v1 := router.Group(basePath)
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}
