package dto

type Credentials struct {
	Password           string `json:"password"`
	Username           string `json:"username"`
	RegistrationNumber string `json:"registration_number"`
}
