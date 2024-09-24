package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update Patient
// @Description Update Patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path string true "Patient identification"
// @Param request body UpdatePatientsRequest true "Request Body"
// @Success 200 {object} UpdatePatientsReponse
// @Failure 400  {object} ErrorResponse
// @Failure 404  {object} ErrorResponse
// @Router /patient/{id} [put]
func UpdatePatients(ctx *gin.Context) {
	// Get the patient ID from the URL parameter
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id-is-required")
		return
	}

	request := UpdatePatientsRequest{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid-request")
	}

	err = request.Validate()
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	patients := schemas.GetSchema()

	if err := db.WithContext(ctx.Request.Context()).First(&patients, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "id-not-found")
		return
	}

	patients = request.UpdatePatientsFieldsChanged(patients)

	if err := db.WithContext(ctx.Request.Context()).Save(&patients).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "InternalServerError")
		return
	}
	sendSucess(ctx, "sucess-update-patient", patients)
}
