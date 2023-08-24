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

type Interests struct {
	Interest   string
	InterestID string
}

type Connection struct {
	UserID          uint
	ConnectedUserID uint
	ConnectionDate  string
}
