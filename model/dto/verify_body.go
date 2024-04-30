package dto

type SuccessLoginBody struct {
	Email       string `json:"userEmail"`
	Password    string `json:"userPassword"`
	AccessToken string `json:"accessToken"`
}

type RequestLoginBody struct {
	Email    string `json:"userEmail"`
	Password string `json:"userPassword"`
}

type VerifyLoginBody struct {
	Email string `json:"userEmail"`
	Otp   string
}

type VerifyLoginBodyResponse struct {
	Token string `json:"token"`
}
