package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show a Patient
// @Description Show a Patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path string true "Patient identification"
// @Success 200 {object} ShowPatientsReponse
// @Failure 400  {object} ErrorResponse
// @Router /patient/{id} [get]
func ShowPatients(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id-is-required")
		return
	}
	patients := schemas.GetSchema()

	if err := db.First(&patients, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "id-not-found")
		return
	}
	sendSucess(ctx, "sucess-find-object", &patients)

}
