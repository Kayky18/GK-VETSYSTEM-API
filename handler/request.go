package handler

import (
	"fmt"

	"github.com/Kayky18/GK_API/schemas"
)

type CreatePatientsRequest struct {
	NameTutor   string  `json:"nameTutor"`
	CPF         string  `json:"cpf"`
	Phone       string  `json:"phone"`
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	Age         int     `json:"age"`
	Weight      float64 `json:"weight"`
	Breed       string  `json:"breed"`
	Temperature float64 `json:"temperature"`
}

type UpdatePatientsRequest struct {
	NameTutor   string  `json:"nameTutor"`
	CPF         string  `json:"cpf"`
	Phone       string  `json:"phone"`
	Name        string  `json:"name"`
	Species     string  `json:"species"`
	Age         int     `json:"age"`
	Weight      float64 `json:"weight"`
	Breed       string  `json:"breed"`
	Temperature float64 `json:"temperature"`
	State       string  `json:"state"`
}

func (r *CreatePatientsRequest) ToSchema() *schemas.Patients {
	return &schemas.Patients{
		NameTutor:   r.NameTutor,
		CPF:         r.CPF,
		Phone:       r.Phone,
		Name:        r.Name,
		Species:     r.Species,
		Age:         r.Age,
		Weight:      r.Weight,
		Breed:       r.Breed,
		Temperature: r.Temperature,
		State:       "Em tratamento",
	}
}

func (request *UpdatePatientsRequest) Validate() error {
	if request.NameTutor != "" || request.Age > 0 || request.Breed != "" || request.CPF != "" || request.Name != "" || request.Phone != "" || request.State != "" || request.Temperature > 0 || request.Weight > 0 || request.Species != "" {
		return nil
	}
	return fmt.Errorf("at least one field must be provided")
}

func (request *CreatePatientsRequest) Validate() error {
	if request.NameTutor != "" && request.Age > 0 && request.CPF != "" && request.Name != "" && request.Phone != "" && request.Species != "" {
		return nil
	}
	return fmt.Errorf("name tutor,species, age, cpf, animal name and phone must be provided. Or some invalid type")
}
