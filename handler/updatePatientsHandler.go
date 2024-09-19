package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

func UpdatePatients(ctx *gin.Context) {
	// Get the patient ID from the URL parameter
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	request := UpdatePatientsRequest{}
	ctx.BindJSON(&request)

	err := request.Validate()
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	patients := schemas.GetSchema()

	if err := db.First(&patients, id).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "id not found")
		return
	}

	if request.NameTutor != "" {
		patients.NameTutor = request.NameTutor
	}
	if request.Age > 0 {
		patients.Age = request.Age
	}
	if request.Breed != "" {
		patients.Breed = request.Breed
	}
	if request.CPF != "" {
		patients.CPF = request.CPF
	}
	if request.Name != "" {
		patients.Name = request.Name
	}
	if request.Phone != "" {
		patients.Phone = request.Phone
	}
	if request.State != "" {
		patients.State = request.State
	}
	if request.Temperature > 0 {
		patients.Temperature = request.Temperature
	}
	if request.Weight > 0 {
		patients.Weight = request.Weight
	}
	if request.Species != "" {
		patients.Species = request.Species
	}

	if err := db.Save(&patients).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "InternalServerError")
		return
	}
	sendSucess(ctx, "sucess-put", patients)
}
