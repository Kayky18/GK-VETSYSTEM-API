package handler

import (
	"fmt"
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

func sendSucess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operation from handler: %s", op),
		"data":    data,
	})
}

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"code":    code,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	Errorcode int    `json:"code"`
}
type CreatePatientsReponse struct {
	Message string                   `json:"message"`
	Data    schemas.PatientsResponse `json:"data"`
}
type DeletePatientsReponse struct {
	Message string                   `json:"message"`
	Data    schemas.PatientsResponse `json:"data"`
}
type ListPatientsReponse struct {
	Message string                     `json:"message"`
	Data    []schemas.PatientsResponse `json:"data"`
}
type BadRequestPatientsReponse struct {
	Message string                   `json:"message"`
	Data    schemas.PatientsResponse `json:"data"`
}
type ShowPatientsReponse struct {
	Message string                   `json:"message"`
	Data    schemas.PatientsResponse `json:"data"`
}
type UpdatePatientsReponse struct {
	Message string                   `json:"message"`
	Data    schemas.PatientsResponse `json:"data"`
}
