package handler

import "github.com/gin-gonic/gin"

func PatientBadRequest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"error": "It is mandatory to provide the patient ID.  Please check the request.Example:(/patient/{ID})."})

}
