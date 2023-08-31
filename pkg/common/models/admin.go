package models

type ForgotPasswordOtpRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

type ForgotPasswordChangePasswordRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ManageUserRequest struct {
	Id string `json:"id"`
}

type AddInterestRequest struct {
	Interest string `json:"interest"`
}
