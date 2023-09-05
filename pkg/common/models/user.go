package models

type RegisterRequestBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CPassword   string `json:"cPassword"`
	PhoneNumber string `json:"phoneNumber"`
	UserName    string `json:"userName"`
	Referral    string `json:"referral"`
	Otp         string `json:"otp"`
}

type OtpValidation struct {
	PhoneNumber string `json:"phoneNumber"`
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	CPassword   string `json:"cPassword"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ValidName struct {
	UserName string `json:"userName"`
}

type ResendOtp struct {
	PhoneNumber string `json:"phoneNumber"`
}

type ForgotPasswordValidateOtpRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Otp         string `json:"otp"`
}

type Interests struct {
	Interest   string
	InterestID string
}

type Connection struct {
	UserID          uint
	ConnectedUserID uint
	ConnectionDate  string
}

type ValidateUserRequest struct {
	Email string `json:"email"`
}

type GoogleLoginRequest struct {
	Token string `json:"token"`
}

type ChangeUserNameRequest struct {
	UserName string `json:"userName"`
}

type ChangeEmailRequest struct {
	Email string `json:"email"`
}

type ChangeEmailVerifyOtpRequest struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type ChangePasswordRequest struct {
	Password string `json:"password"`
}

type ChangePhoneNumberOtpRequest struct {
	PhoneNumber string `json:"phoneNumber"`
}

type ChangePhoneNumberRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Otp         string `json:"otp"`
}
