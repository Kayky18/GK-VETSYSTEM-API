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
// @Param id path string true "Patient identification"
// @Success 200 {object} ShowPatientsReponse
// @Failure 400  {object} ErrorResponse
// @Router /patient/{id} [get]
func ShowPatients(ctx *gin.Context) {
	id := ctx.Param("id")

	patients := schemas.GetSchema()

	if err := db.First(&patients, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSucess(ctx, "sucess-find-object", &patients)

}
