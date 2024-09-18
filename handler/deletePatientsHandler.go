package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

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
