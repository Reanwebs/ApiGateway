package models

type RegisterRequestBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	UserName    string `json:"userName"`
	Referral    string `json:"referral"`

}

type OtpValidation struct {
	Otp         string `json:"otp"`
	PhoneNumber string `json:"phoneNumber"`
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
