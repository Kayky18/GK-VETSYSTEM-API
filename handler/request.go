package handler

import (
	"fmt"
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
	return fmt.Errorf("name tutor, age, cpf, animal name and phone must be provided")
}
