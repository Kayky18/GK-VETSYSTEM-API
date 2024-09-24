package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Patients
// @Description List Patients
// @Tags Patients
// @Accept json
// @Produce json
// @Success 200 {object} ListPatientsReponse
// @Failure 500  {object} ErrorResponse
// @Router /patients [get]
func ListPatients(ctx *gin.Context) {
	// Get the patient list from the database
	patients := schemas.GetListSchema()
	// Return the patient list in the response
	if err := db.Find(&patients).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error-to-list-patients")
		return
	}
	sendSucess(ctx, "sucess-to-list", &patients)
}
