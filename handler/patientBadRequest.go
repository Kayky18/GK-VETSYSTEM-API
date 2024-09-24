package handler

import "github.com/gin-gonic/gin"

// @BasePath /api/v1

// @Summary Bad Request Patient
// @Description Bad Request Patient
// @Tags Patients
// @Accept json
// @Produce json
// @Success 200 {object} BadRequestPatientsReponse
// @Router /patient [get]
func PatientBadRequest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"error": "It is mandatory to provide the patient ID.  Please check the request.Example:(/patient/{ID})."})
}
