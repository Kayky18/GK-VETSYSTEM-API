package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create Patient
// @Description Create a new Patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param request body CreatePatientsRequest true "Request Body"
// @Success 200 {object} CreatePatientsReponse
// @Failure 400  {object} ErrorResponse
// @Failure 500  {object} ErrorResponse
// @Router /patient/create [post]
func CreatePatients(ctx *gin.Context) {
	request := CreatePatientsRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error-to-bind-json")
	}

	if err = request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, "request-validate-error")
		return
	}
	//Send to Schema
	patients := request.ToSchema()

	if err := db.Create(&patients).Error; err != nil {
		logger.Errorf("Error creating patient in database: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "database-creating-error")
		return
	}
	sendSucess(ctx, "sucess-create-patient", &patients)
}
