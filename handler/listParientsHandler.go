package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

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
