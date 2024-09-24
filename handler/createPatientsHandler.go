package handler

import (
	"net/http"

	"github.com/Kayky18/GK_API/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create Patient
// @Description Create a new Patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param request body CreatePatientsRequest true "Request Body"
// @Success 200 {object} CreatePatientsReponse
// @Failure 400  {object} ErrorResponse
// @Failure 500  {object} ErrorResponse
// @Router /patient/create [post]
func CreatePatients(ctx *gin.Context) {
	request := CreatePatientsRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	patients := schemas.GetSchema()

	// Create a new patient
	patients.NameTutor = request.NameTutor
	patients.CPF = request.CPF
	patients.Phone = request.Phone
	patients.Name = request.Name
	patients.Age = request.Age
	patients.Weight = request.Weight
	patients.Breed = request.Breed
	patients.Temperature = request.Temperature
	patients.State = "Em tratamento"
	patients.Species = request.Species

	if err := db.Create(&patients).Error; err != nil {
		logger.Errorf("ERROR CREATING DATABASE: %v", err)
		sendError(ctx, http.StatusInternalServerError, "creating-error")
		return
	}
	sendSucess(ctx, "sucess-create", &patients)
}
