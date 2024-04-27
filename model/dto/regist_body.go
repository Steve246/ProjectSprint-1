package dto

type RequestRegistBody struct {
	Name     string `json:"userName"`
	Email    string `json:"userEmail"`
	Password string `json:"userPassword"`
}

type VerifyRegistBody struct {
	Otp string
}
