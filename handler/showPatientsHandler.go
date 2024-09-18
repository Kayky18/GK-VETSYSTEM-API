package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

func ShowPatients(ctx *gin.Context) {
	id := ctx.Param("id")

	patients := schemas.GetSchema()

	if err := db.First(&patients, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	sendSucess(ctx, "sucess-find-object", &patients)

}
