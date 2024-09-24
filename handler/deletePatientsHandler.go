package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete Patient
// @Description Delete a Patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path string true "Patient identification"
// @Success 200 {object} DeletePatientsReponse
// @Failure 404  {object} ErrorResponse
// @Router /patient/{id} [delete]
func DeletePatients(ctx *gin.Context) {
	// Get the patient ID from the URL parameter
	id := ctx.Param("id")
	// Delete the patient from the database
	patients := schemas.GetSchema()

	//To have the data deleted
	if err := db.First(&patients, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "Error to find opening with id")
		return
	}

	//Deleted Data
	if err := db.Delete(&patients).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error-to-delete-patient")
		return
	}

	sendSucess(ctx, "sucess-delete", &patients)
}
