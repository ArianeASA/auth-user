package dto

type User struct {
	Username           string `json:"username"`
	RegistrationNumber string `json:"registration_number"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	Password           string `json:"password"`
}
