package models

type RegisterRequestBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	UserName    string `json:"userName"`
}

type OtpValidation struct {
	Otp         string `json:"otp"`
	PhoneNumber string `json:"phoneNumber"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
