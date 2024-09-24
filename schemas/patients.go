package schemas

import "gorm.io/gorm"

type Patients struct {
	*gorm.Model
	NameTutor   string  //Name Tutor
	CPF         string  //CPF Tutor
	Phone       string  //Phone  Tutor
	Name        string  //Name Animal
	Species     string  //Species
	Age         int     //Age Animal
	Weight      float64 //Weight Animal
	Breed       string  //Breed
	Temperature float64 //Temperature Animal
	State       string  //State in clinic

}
type PatientsResponse struct {
	NameTutor string  `json:"nameTutor"`
	CPF       string  `json:"cpf"`
	Phone     string  `json:"phone"`
	Name      string  `json:"name"`
	Species   string  `json:"species"`
	Age       int     `json:"age"`
	Weight    float64 `json:"weight"`
	Breed     string  `json:"breed"`
	State     string  `json:"state"`
}

func GetSchema() *Patients {
	return &Patients{}
}

func GetListSchema() []*Patients {
	return []*Patients{}
}
